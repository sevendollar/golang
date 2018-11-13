package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "net/url"
        "os"
        "path"
        "runtime"
        "strings"
        "sync"

        "github.com/PuerkitoBio/goquery"
)

var (
        wg          = sync.WaitGroup{}
        downloadDir = "/tmp/pdf/"
        baseUrl     = "http://library.globalchalet.net/Authors/"
        //baseUrl     = "http://library.globalchalet.net/Authors/Stine,%20R%20L/"
        //baseUrl   = "http://library.globalchalet.net/Authors/Blizzard%20Collection/"
        fileQueue = make(chan string)
        urlQueue  = make(chan string)
)

func getUrl(target string, fileQueue chan string, urlQueue chan string) {
        defer wg.Done()
        res, err := http.Get(target)
        if err != nil {
                log.Fatal(err)
                return
        }
        defer res.Body.Close()

        doc, err := goquery.NewDocumentFromReader(res.Body)
        if err != nil {
                log.Fatal(err)
                return
        }

        doc.Find("a").Each(func(i int, s *goquery.Selection) {
                if href, ok := s.Attr("href"); ok {
                        switch {
                        case href == "/Authors/" || strings.HasPrefix(href, "http"):
                                //log.Println("invaild path:", href)
                        case strings.HasSuffix(href, ".pdf") || strings.HasSuffix(href, ".mp3"):
                                fileQueue <- target + href
                        case strings.HasSuffix(href, "/"):
                                urlQueue <- href
                        }
                }
        })
        //close(fileQueue)
}

func getFile(fileUrl string) {
        defer wg.Done()
        res, err := http.Get(fileUrl)
        if err != nil {
                log.Fatal(err)
                return
        }
        defer res.Body.Close()

        data, err := ioutil.ReadAll(res.Body)
        if err != nil {
                log.Fatal(err)
                return
        }

        base, err := url.QueryUnescape(path.Base(fileUrl))
        if err != nil {
                log.Fatal(err)
                return
        }

        subDir, err := url.QueryUnescape(path.Base(path.Dir(fileUrl)))
        if err != nil {
                log.Fatal(err)
                return
        }

        err = mkdir(downloadDir + subDir)
        if err != nil {
                log.Fatal(err)
                return
        }

        err = ioutil.WriteFile(downloadDir+subDir+"/"+base, data, 0644)
        if err != nil {
                log.Fatal(err)
                return
        }
        fmt.Printf("writing file... %v\n", base)
}

func mkdir(dir string) error {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
                err = os.MkdirAll(dir, 0755)
                if err != nil {
                        return err
                }
        }
        return nil
}

func index(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
        <html>
          <body>
          pdf collection
          <a href="/file">go</a>
          </body>
        </html>
        `))
}

func init() {
        mkdir(downloadDir)
}

func main() {
        runtime.GOMAXPROCS(8)

        wg.Add(1)
        go getUrl(baseUrl, fileQueue, urlQueue)

        wg.Add(1)
        go func() {
                defer wg.Done()
                for {
                        select {
                        case v, ok := <-urlQueue:
                                if ok {
                                        wg.Add(1)
                                        go getUrl(baseUrl+v, fileQueue, urlQueue)
                                } else {
                                        urlQueue = nil
                                }
                        }
                        if urlQueue == nil {
                                break
                        }
                }
        }()

        wg.Add(1)
        go func() {
                defer wg.Done()
                for {
                        select {
                        case v, ok := <-fileQueue:
                                if ok {
                                        wg.Add(1)
                                        go getFile(v)
                                } else {
                                        fileQueue = nil
                                }
                        }

                        if fileQueue == nil {
                                break
                        }

                }
        }()

        http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(downloadDir))))
        http.HandleFunc("/", index)

        fmt.Println("serving... ")
        http.ListenAndServe(":80", nil)

        wg.Wait()

}

