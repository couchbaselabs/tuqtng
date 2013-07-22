package main

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

func process(cache string, src string, f os.FileInfo, err error) error {
	if strings.HasPrefix(f.Name(), ".") {
		return nil
	}

	dst := cache + src

	if f.IsDir() {
		if err := os.MkdirAll(dst, 0700); err != nil {
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
		log.Println("Expanding", f.Name())
		dst = strings.TrimSuffix(dst, ".md") + ".html"
		dbuf = blackfriday.MarkdownCommon(sbuf)
	}

	if err := ioutil.WriteFile(dst, dbuf, f.Mode()); err != nil {
		return err
	}

	return nil
}

func main() {
	tmpDir, _ := ioutil.TempDir("", "tut")
	tmpDir += string(os.PathSeparator)
	defer os.RemoveAll(tmpDir)
	log.Println("Workdir", tmpDir)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.RemoveAll(tmpDir)
		log.Fatalln("Stopped")
	}()

	walker := func(src string, f os.FileInfo, err error) error {
		return process(tmpDir, src, f, err)
	}

	if err := filepath.Walk(".", walker); err != nil {
		log.Fatalln("Filewalk", err)
	}

	url, _ := url.Parse("http://localhost:8093")
	rp := httputil.NewSingleHostReverseProxy(url)
	http.Handle("/query", rp)

	fs := http.FileServer(http.Dir("./")) // make this into tempDir
	http.Handle("/tutorial/", http.StripPrefix("/tutorial/", fs))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
