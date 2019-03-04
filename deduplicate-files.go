package main

import (
        "crypto/sha512"
        "fmt"
        "io/ioutil"
        "os"
        "path/filepath"
)

var (
        files = make(map[[sha512.Size]byte]string)
        arg   = ""
)

func main() {

        err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
                if !info.IsDir() {
                        data, err := ioutil.ReadFile(path)
                        if err != nil {
                                fmt.Println("Error:", err)
                                return err
                        }
                        hash := sha512.Sum512(data)
                        if v, ok := files[hash]; ok {
                                fmt.Printf("%v is a duplicate of %v\n", path, v)
                                if len(os.Args) >= 2 && os.Args[1] == "--delete" {
                                        if err = os.Remove(path); err != nil {
                                                fmt.Println("Error:", err)
                                                return err
                                        } else {
                                                fmt.Printf("%v deleted...\n", path)
                                        }
                                }
                        } else {
                                files[hash] = path
                        }
                }
                return nil
        })
        if err != nil {
                fmt.Println("Error:", err)
                return
        }
        fmt.Println("Scan completed!")

}

