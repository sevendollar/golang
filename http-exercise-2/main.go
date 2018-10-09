package main

import (
        "fmt"
        "net/http"
        "os"
        "strings"
)

func checkErr(err error) {
        if err != nil {
                fmt.Println("Error:", err)
                os.Exit(2)
        }
}

func getCookie(cookies string) (m map[string]string) {
        m = make(map[string]string)
        for _, cookie := range strings.Split(cookies, "; ") {
                v := strings.Split(cookie, "=")
                if len(v) == 2 {
                        m[v[0]] = v[1]
                } else {
                        return
                }
        }
        return
}

func main() {
        server := &http.Server{
                Addr: ":80",
        }

        http.HandleFunc("/", index)

        http.HandleFunc("/signup/", signup)
        http.HandleFunc("/signup", signup)
        http.HandleFunc("/process_signup/", process_signup)
        http.HandleFunc("/process_signup", process_signup)

        http.HandleFunc("/signin/", signin)
        http.HandleFunc("/signin", signin)
        http.HandleFunc("/autheticate/", autheticate)
        http.HandleFunc("/autheticate", autheticate)

        http.HandleFunc("/signout/", signout)
        http.HandleFunc("/signout", signout)

        http.HandleFunc("/verify/", verify)
        http.HandleFunc("/verify", verify)

        fmt.Println("Serving...")
        server.ListenAndServe()
}

