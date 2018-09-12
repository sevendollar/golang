package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
)

type user_profile struct {
    Id        int
    Name      string
    Ocupation string
}

func main() {
    var user_profiles []*user_profile

    p1 := []*user_profile{
        &user_profile{Id: 1, Name: "jef", Ocupation: "golang developer"},
        &user_profile{Id: 2, Name: "joy", Ocupation: "erp consultant"},
        &user_profile{Id: 3, Name: "rex", Ocupation: "general manager"},
    }

    fw, err := os.Create("test.csv")
    if err != nil {
        panic(err)
    }
    defer fw.Close()
    w := csv.NewWriter(fw)
    for _, user := range p1 {
        w.Write([]string{strconv.Itoa(user.Id), user.Name, user.Ocupation})
    }
    w.Flush()

    fo, err := os.Open("test.csv")
    if err != nil {
        panic(err)
    }
    defer fo.Close()

    r := csv.NewReader(fo)
    data, err := r.ReadAll()
    if err != nil {
        panic(err)
    }

    for _, line := range data {
        id, _ := strconv.Atoi(line[0])
        user_profiles = append(user_profiles, &user_profile{
            Id:        id,
            Name:      line[1],
            Ocupation: line[2],
        })
    }

    user_profiles_json, err := json.MarshalIndent(user_profiles, "", "  ")
    if err != nil {
        panic(err)
    }

    jsonFile, err := os.Create("json.text")
    if err != nil {
        panic(err)
    }
    defer jsonFile.Close()
    jsonFile.Write(user_profiles_json)
    jsonFile.Write([]byte("\n"))

    fmt.Println("created csv file, read from it and then convert into json format....")
    fmt.Println()
    fmt.Println(string(user_profiles_json))
}
