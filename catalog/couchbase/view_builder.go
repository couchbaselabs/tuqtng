package couchbase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
	"strings"
	"time"

	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
)

type ddocJSON struct {
	cb.DDoc
	IndexOn       []string `json:"indexOn"`
	IndexChecksum int      `json:"indexChecksum"`
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

	err = inst.WaitForIndex()
	if err != nil {
		return nil, err
	}

	return &inst, nil
}

func (vi *viewIndex) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("name: %v ", vi.name))
	buf.WriteString(fmt.Sprintf("on: %v ", vi.on))
	buf.WriteString(fmt.Sprintf("using: %v ", vi.using))
	buf.WriteString(fmt.Sprintf("ddoc: %v ", *vi.ddoc))
	buf.WriteString(fmt.Sprintf("bucket: %v ", *vi.bucket))
	return buf.String()
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

func loadViewIndexes(b *bucket) ([]*catalog.Index, error) {

	rows, err := b.cbbucket.GetDDocs()
	if err != nil {
		return nil, err
	}

	inames := make([]string, 0, len(rows.Rows))
	for _, row := range rows.Rows {
		cdoc := row.DDoc
		id := cdoc.Meta["id"].(string)
		if !strings.HasPrefix(id, "_design/ddl_") {
			continue
		}
		iname := strings.TrimPrefix(id, "_design/ddl_")
		inames = append(inames, iname)
	}

	indexes := make([]*catalog.Index, 0, len(inames))
	for _, iname := range inames {
		ddname := "ddl_" + iname
		jdoc, err := getDesignDoc(b, ddname)
		if err != nil {
			return nil, err
		}
		jview, ok := jdoc.Views[iname]
		if !ok {
			return nil, errors.New("Missing view for index " + iname)
		}

		exprlist := make([]ast.Expression, 0, len(jdoc.IndexOn))

		for _, ser := range jdoc.IndexOn {
			// HACK - remove this when Unmarshall supports META()
			if iname == PRIMARY_INDEX {
				meta := ast.NewFunctionCall("meta", ast.FunctionArgExpressionList{})
				mdid := ast.NewDotMemberOperator(meta, ast.NewProperty("id"))
				exprlist = append(exprlist, mdid)
			} else {
				expr, err := ast.UnmarshalExpression([]byte(ser))
				if err != nil {
					return nil, errors.New("Cannot unmarshal expression for index " + iname)
				}
				exprlist = append(exprlist, expr)
			}
		}
		if len(exprlist) != len(jdoc.IndexOn) {
			continue
		}

		ddoc := designdoc{
			name:     ddname,
			viewname: iname,
			mapfn:    jview.Map,
			reducefn: jview.Reduce,
		}
		if ddoc.checksum() != jdoc.IndexChecksum {
			return nil, errors.New("Warning - checksum failed on index " + iname)
		}

		var index catalog.Index

		if iname == PRIMARY_INDEX {
			index = &primaryIndex{
				viewIndex{
					name:   iname,
					bucket: b,
					using:  catalog.VIEW,
					ddoc:   &ddoc,
					on:     exprlist,
				},
			}
			indexes = append(indexes, &index)
		} else {
			index = &viewIndex{
				name:   iname,
				bucket: b,
				using:  catalog.VIEW,
				ddoc:   &ddoc,
				on:     exprlist,
			}
			indexes = append(indexes, &index)
		}
	}

	return indexes, nil
}

func newPrimaryIndex(b *bucket) (*primaryIndex, error) {
	ddoc := newPrimaryDDoc()
	meta := ast.NewFunctionCall("meta", ast.FunctionArgExpressionList{})
	mdid := ast.NewDotMemberOperator(meta, ast.NewProperty("id"))
	inst := primaryIndex{
		viewIndex{
			name:   PRIMARY_INDEX,
			using:  catalog.VIEW,
			on:     catalog.IndexKey{mdid},
			ddoc:   ddoc,
			bucket: b,
		},
	}

	err := inst.putDesignDoc()
	if err != nil {
		return nil, err
	}

	err = inst.WaitForIndex()
	if err != nil {
		return nil, err
	}

	return &inst, nil
}

func newAllDocsIndex(b *bucket) *primaryIndex {
	meta := ast.NewFunctionCall("meta", ast.FunctionArgExpressionList{})
	mdid := ast.NewDotMemberOperator(meta, ast.NewProperty("id"))
	ddoc := designdoc{name: "", viewname: "_all_docs"}
	idx := primaryIndex{
		viewIndex{
			name:   ALLDOCS_INDEX,
			using:  catalog.VIEW,
			on:     catalog.IndexKey{mdid},
			ddoc:   &ddoc,
			bucket: b,
		},
	}
	return &idx
}

func newPrimaryDDoc() *designdoc {
	var doc designdoc
	line := strings.Replace(templPrimary, "$rnd", strconv.Itoa(int(rand.Int31())), -1)
	line = strings.Replace(line, "$string", strconv.Itoa(TYPE_STRING), -1)
	doc.mapfn = line
	doc.reducefn = ""
	doc.name = "ddl_" + PRIMARY_INDEX
	doc.viewname = PRIMARY_INDEX
	return &doc
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

	line = strings.Replace(templEnd, "$rnd", strconv.Itoa(int(rand.Int31())), -1)
	fmt.Fprint(buf, line)

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
	put.IndexChecksum = idx.ddoc.checksum()

	put.IndexOn = make([]string, len(idx.on))
	for idx, expr := range idx.on {
		ser, err := json.Marshal(expr)
		if err != nil {
			return err
		}
		put.IndexOn[idx] = string(ser)
	}

	if err := idx.bucket.cbbucket.PutDDoc(idx.DDocName(), &put); err != nil {
		return err
	}

	var saved *ddocJSON = nil
	var err error = nil

	// give the PUT some time to register
	for i := 0; i < 3; i++ {
		if i > 1 {
			time.Sleep(time.Duration(i*3) * time.Second)
		}

		saved, err = getDesignDoc(idx.bucket, idx.DDocName())
		if err == nil {
			break
		}
	}

	if err != nil {
		return errors.New("Creating index '" + idx.name + "' failed: " + err.Error())
	}

	if saved.IndexChecksum != idx.ddoc.checksum() {
		return errors.New("Checksum mismatch after creating index '" + idx.name + "'")
	}

	return nil
}

func (ddoc *designdoc) checksum() int {
	mapSum := crc32.ChecksumIEEE([]byte(ddoc.mapfn))
	reduceSum := crc32.ChecksumIEEE([]byte(ddoc.reducefn))
	return int(mapSum + reduceSum)
}

func getDesignDoc(b *bucket, ddocname string) (*ddocJSON, error) {
	var ddoc ddocJSON
	err := b.cbbucket.GetDDoc(ddocname, &ddoc)
	if err != nil {
		return nil, err
	}
	return &ddoc, nil
}

func (idx *viewIndex) DropViewIndex() error {
	if err := idx.bucket.cbbucket.DeleteDDoc(idx.ddoc.name); err != nil {
		return err
	}
	return nil
}

func (idx *viewIndex) WaitForIndex() error {
	var err error
	// if we have got this far, very likely any errors are
	// due to index not yet being noticed by the system.
	for i := 0; i < 3; i++ {
		if i > 1 {
			time.Sleep(time.Duration(i*3) * time.Second)
		}
		_, err = idx.bucket.cbbucket.View(
			idx.ddoc.name,
			idx.ddoc.viewname,
			map[string]interface{}{
				"start_key": []interface{}{"thing"},
				"end_key":   []interface{}{"thing", map[string]string{}},
				"stale":     false,
			})
		if err == nil {
			break
		}
	}
	return err
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
		return e, errors.New("Expression is not supported by indexing currently: " + e.String())

	}
	return e, nil
}
