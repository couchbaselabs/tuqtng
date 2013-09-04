package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/couchbaselabs/clog"
	"github.com/russross/blackfriday"
)

type Index struct {
	Entries []string `json:"entries`
}

func setup(cacheDir string, src string, srcF os.FileInfo, idx *Index, err error) error {
	if strings.HasPrefix(srcF.Name(), ".") || strings.Contains(src, "/.") {
		return nil
	}

	dst := cacheDir + src
	if strings.HasSuffix(dst, ".md") {
		dst = strings.TrimSuffix(dst, ".md") + ".html"
	}
	dstF, dstE := os.Stat(dst)

	// if up to date, skip
	if !os.IsNotExist(dstE) &&
		(srcF.IsDir() || dstF.ModTime().After(srcF.ModTime())) {
		return nil
	}

	// copy to cache
	if os.IsNotExist(dstE) {
		clog.Log("Copying %s", srcF.Name())
	} else {
		clog.Log("Updating %s", srcF.Name())
	}

	if srcF.IsDir() {
		if err := os.MkdirAll(dst, srcF.Mode()); err != nil {
			return err
		}
		return nil
	}

	sbuf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	dbuf := sbuf
	if strings.HasSuffix(src, ".md") {
		dbuf = blackfriday.MarkdownCommon(sbuf)
	}
	if err := ioutil.WriteFile(dst, dbuf, srcF.Mode()); err != nil {
		return err
	}

	headrx := regexp.MustCompile("##[^\n]+\n")
	found := headrx.FindString(string(sbuf))
	if len(found) > 0 {
		found = strings.TrimSpace(strings.TrimPrefix(found, "##"))
		idx.Entries = append(idx.Entries, found)
	}

	return nil
}

func main() {
	tempDir, _ := ioutil.TempDir("", "tut")
	tempDir += string(os.PathSeparator)
	defer os.RemoveAll(tempDir)
	clog.Log("Workdir %s", tempDir)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.RemoveAll(tempDir)
		clog.Fatal("Stopped")
	}()

	var idx Index
	walker := func(src string, f os.FileInfo, err error) error {
		return setup(tempDir, src, f, &idx, err)
	}

	idx.Entries = []string{}
	if err := filepath.Walk("./content/", walker); err != nil {
		clog.Fatalf("Filewalk %v", err)
	}

	http.Handle("/", http.RedirectHandler("/tutorial/tutorial.html#1", 302))

	url, _ := url.Parse("http://localhost:8093")
	rp := httputil.NewSingleHostReverseProxy(url)
	http.Handle("/query", rp)

	fs := http.FileServer(http.Dir(tempDir + "/content/"))
	http.Handle("/tutorial/", http.StripPrefix("/tutorial/", fs))

	getindex := func(w http.ResponseWriter, r *http.Request) {
		json, err := json.Marshal(idx)
		if err != nil {
			clog.Fatalf("During Index JSON Marshal %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
	http.HandleFunc("/tutorial/index.json", getindex)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		clog.Fatalf("ListenAndServe %v", err)
	}

	clog.Log("Running at http:///localhost:8000/")
	go func() {
		for {
			idx.Entries = []string{}
			filepath.Walk("./content/", walker)
			time.Sleep(2 * time.Second)
		}
	}()

}
