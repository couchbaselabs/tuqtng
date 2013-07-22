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
	"time"
)

func setup(cacheDir string, src string, srcF os.FileInfo, err error) error {
	if strings.HasPrefix(srcF.Name(), ".") {
		return nil
	}

	dst := cacheDir + src
	if strings.HasSuffix(dst, ".md") {
		dst = strings.TrimSuffix(dst, ".md") + ".html"
	}
	dstF, _ := os.Stat(dst)

	// if up to date, skip
	if srcF.IsDir() && dstF != nil {
		return nil
	}
	if dstF != nil && dstF.ModTime().After(srcF.ModTime()) {
		return nil
	}

	// copy to cache
	if srcF.IsDir() {
		log.Println("Creating Dir", srcF.Name())
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
	log.Println("Copying", srcF.Name())
	if err := ioutil.WriteFile(dst, dbuf, srcF.Mode()); err != nil {
		return err
	}
	return nil
}

func main() {
	tempDir, _ := ioutil.TempDir("", "tut")
	tempDir += string(os.PathSeparator)
	defer os.RemoveAll(tempDir)
	log.Println("Workdir", tempDir)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.RemoveAll(tempDir)
		log.Fatalln("Stopped")
	}()

	walker := func(src string, f os.FileInfo, err error) error {
		return setup(tempDir, src, f, err)
	}

	if err := filepath.Walk("./content/", walker); err != nil {
		log.Fatalln("Filewalk", err)
	}

	go func() {
		for {
			filepath.Walk("./content/", walker)
			time.Sleep(3 * time.Second)
		}
	}()

	http.Handle("/", http.RedirectHandler("/tutorial/tutorial.html", 302))

	url, _ := url.Parse("http://localhost:8093")
	rp := httputil.NewSingleHostReverseProxy(url)
	http.Handle("/query", rp)

	fs := http.FileServer(http.Dir(tempDir + "/content/"))
	http.Handle("/tutorial/", http.StripPrefix("/tutorial/", fs))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
