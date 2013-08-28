package couchbase

import (
	"bytes"
	"fmt"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/indexing/api"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"hash/crc32"
	"strconv"
	"strings"
)

type ddocJSON struct {
	cb.DDocJSON
	IndexOn       string `json:"indexOn"`
	IndexChecksum int    `json:"indexChecksum"`
}

func newViewIndex(name string, on catalog.IndexKey, bkt *bucket) (*viewIndex, error) {

	doc, err := newDesignDoc(name, on)
	if err != nil {
		return nil, err
	}

	inst := viewIndex{
		name:   name,
		using:  catalog.VIEW,
		on:     on,
		ddoc:   doc,
		bucket: bkt,
	}

	err = inst.putDesignDoc()
	if err != nil {
		return nil, err
	}

	return &inst, nil
}

func newDesignDoc(idxname string, on catalog.IndexKey) (*designdoc, error) {
	var doc designdoc

	doc.name = "ddl_" + idxname
	doc.viewname = idxname

	err := generateMap(on, &doc)
	if err != nil {
		return nil, err
	}

	err = generateReduce(on, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func newPrimaryIndex(b *bucket, ddname string, view string) (*primaryIndex, error) {
	meta := ast.NewFunctionCall("meta", ast.FunctionArgExpressionList{})
	mdid := ast.NewDotMemberOperator(meta, ast.NewProperty("id"))
	name := "#primary"
	ddoc := designdoc{name: ddname, viewname: view}
	idx := primaryIndex{
		viewIndex{
			name:   name,
			using:  catalog.PRIMARY,
			on:     catalog.IndexKey{mdid},
			ddoc:   &ddoc,
			bucket: b,
		},
	}
	return &idx, nil
}

func generateMap(on catalog.IndexKey, doc *designdoc) error {

	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, templStart)
	line := strings.Replace(templFunctions, "$null", strconv.Itoa(TYPE_NULL), -1)
	line = strings.Replace(line, "$boolean", strconv.Itoa(TYPE_BOOLEAN), -1)
	line = strings.Replace(line, "$number", strconv.Itoa(TYPE_NUMBER), -1)
	line = strings.Replace(line, "$string", strconv.Itoa(TYPE_STRING), -1)
	line = strings.Replace(line, "$array", strconv.Itoa(TYPE_ARRAY), -1)
	line = strings.Replace(line, "$object", strconv.Itoa(TYPE_OBJECT), -1)
	fmt.Fprintln(buf, line)

	keylist := new(bytes.Buffer)
	for idx, expr := range on {

		walker := NewWalker()
		_, err := walker.Visit(expr)
		if err != nil {
			return err
		}

		jvar := fmt.Sprintf("key%v", idx+1)
		line := strings.Replace(templExpr, "$var", jvar, -1)
		line = strings.Replace(line, "$path", walker.JS(), -1)
		fmt.Fprint(buf, line)

		if idx > 0 {
			fmt.Fprint(keylist, ", ")
		}
		fmt.Fprint(keylist, jvar)
	}

	line = strings.Replace(templKey, "$keylist", keylist.String(), -1)

	fmt.Fprint(buf, line)
	fmt.Fprint(buf, templEmit)
	fmt.Fprint(buf, templEnd)
	doc.mapfn = buf.String()

	return nil
}

func generateReduce(on catalog.IndexKey, doc *designdoc) error {
	doc.reducefn = ""
	return nil
}

func (idx *viewIndex) putDesignDoc() error {
	var view cb.ViewDefinition
	view.Map = idx.ddoc.mapfn

	var put ddocJSON
	put.Views = make(map[string]cb.ViewDefinition)
	put.Views[idx.name] = view
	put.IndexChecksum = checksum(idx.ddoc)

	if err := idx.bucket.cbbucket.PutDDoc(idx.DDocName(), &put); err != nil {
		return err
	}

	err := idx.checkDesignDoc()
	if err != nil {
		return api.DDocCreateFailed
	}

	return nil
}

func checksum(ddoc *designdoc) int {
	mapSum := crc32.ChecksumIEEE([]byte(ddoc.mapfn))
	reduceSum := crc32.ChecksumIEEE([]byte(ddoc.reducefn))
	return int(mapSum + reduceSum)
}

func (idx *viewIndex) checkDesignDoc() error {
	var ddoc ddocJSON

	if err := idx.bucket.cbbucket.GetDDoc(idx.DDocName(), &ddoc); err != nil {
		return err
	}

	if ddoc.IndexChecksum != checksum(idx.ddoc) {
		return api.DDocChanged
	}

	return nil
}

func (idx *viewIndex) DropViewIndex() error {
	if err := idx.bucket.cbbucket.DeleteDDoc(idx.DDocName()); err != nil {
		return err
	}
	return nil
}

// AST to JS conversion
type JsStatement struct {
	js bytes.Buffer
}

func NewWalker() *JsStatement {
	var js JsStatement
	return &js
}

func (this *JsStatement) JS() string {
	return this.js.String()
}

// inorder traversal of the AST to get JS expression out of it
func (this *JsStatement) Visit(e ast.Expression) (ast.Expression, error) {
	switch expr := e.(type) {

	case *ast.DotMemberOperator:
		if this.js.Len() == 0 {
			this.js.WriteString("doc.")
		}
		_, err := expr.Left.Accept(this)
		if err != nil {
			return nil, err
		}
		this.js.WriteString(".")
		_, err = expr.Right.Accept(this)
		if err != nil {
			return nil, err
		}

	case *ast.BracketMemberOperator:
		if this.js.Len() == 0 {
			this.js.WriteString("doc.")
		}
		_, err := expr.Left.Accept(this)
		if err != nil {
			return nil, err
		}
		this.js.WriteString("[")
		_, err = expr.Right.Accept(this)
		if err != nil {
			return nil, err
		}
		this.js.WriteString("]")

	case *ast.Property:
		if this.js.Len() == 0 {
			this.js.WriteString("doc.")
		}
		this.js.WriteString(expr.Path)

	case *ast.LiteralNumber:
		this.js.WriteString(fmt.Sprintf("%v", expr.Val))

	case *ast.LiteralString:
		this.js.WriteString(expr.Val)

	default:
		return e, api.ExprNotSupported

	}
	return e, nil
}
