package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	jef "net/url"
	"path"
	"strings"
	"sync"
	"time"
)

const (
	url  = "http://library.globalchalet.net/Authors/Stine,%20R%20L/"
	dir  = "/tmp"
	port = "80"
)

var (
	wg sync.WaitGroup
)

func getHref(url string) (hrefs []string, err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("href"); ok {
			if strings.Contains(v, ".pdf") {
				hrefs = append(hrefs, fmt.Sprint(url, v))
			}
		}
	})
	return hrefs, nil
}

func downloader(url string, dir string) (err error) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	filename, err := jef.QueryUnescape(url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = ioutil.WriteFile(dir+path.Base(filename), data, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//fmt.Println("succeed!")
	return nil
}

func downloaderWraper() {
	hrefs, err := getHref(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(hrefs) == 0 {
		fmt.Println("there has no hrefs found in here.")
		return
	} else {
		wg.Add(len(hrefs))
		for _, v := range hrefs {
			//download href
			go downloader(v, dir)
		}
	}
}

func fileChecker() {
	defer wg.Done()

	for {
		fileExist := false
		files, err := ioutil.ReadDir(dri)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(files) != 0 {
			for _, v := range files {
				fileExist = fileExist || strings.Contains(v.Name(), ".pdf")
			}
		}

		fmt.Println(fileExist, time.Now())

		if !fileExist {
			downloaderWraper()
			fmt.Println("getting file...")
		}
		time.Sleep(time.Duration(30) * time.Second)
	}
	fmt.Println("exit fileChecker()")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
        <html>
          <body>
        jef's pdf collection
                <a href="/file">go</a>
          </body>
        </html>
        `))
}

func main() {

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/", index)
	mux.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(dir))))

	//downloaderWraper()

	wg.Add(1)
	go fileChecker()

	fmt.Println("web serving...")
	server.ListenAndServe()

	wg.Wait()
}
