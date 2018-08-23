package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "math/rand"
    "os"
    "time"
)

var (
    usernames = []string{
        "Wil van der Aalst",
        "Scott Aaronson",
        "Hal Abelson",
        "Samson Abramsky",
        "Andrew Appel",
        "Andrew Barto",
        "Peter Bernus",
        "Ken Thompson",
        "Daniel J. Bernstein",
        "Anita Borg",
        "Kurt Bollacker",
        "Lenore Blum",
        "Sue Black",
    }
    skills = []string{
        "python",
        "java",
        "javascript",
        "php",
        "c",
        "c++",
        "c#",
        "golang",
        "ruby",
        "shell",
    }
)

type User struct {
    ID          bson.ObjectId `bson:"_id,omitempty"`
    Username    string        `json:"username"`
    Age         int           `json:"age"`
    Sex         int           `json:"sex"`
    Skill       []string      `json:"skill"`
    CreatedTime time.Time     `json:"created_time"`
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func shuffle(x []string) []string {
    for i := range x {
        j := rand.Intn(i + 1)
        x[i], x[j] = x[j], x[i]
    }
    return x
}

func main() {
    count := flag.Int("c", 5, "how many times you want to insert the datasets.")
    mongo_dial_info := &mgo.DialInfo{
        Addrs:    []string{"10.5.1.160:27017"},
        Timeout:  3 * time.Second,
        Database: "admin",
        Username: "jef",
        Password: "mongopw",
    }
    tmp_user := User{}
    tmp_info := []User{}

    session, err := mgo.DialWithInfo(mongo_dial_info)
    if err != nil {
        log.Println("Error", err)
        os.Exit(2)
    }
    defer session.Close()
    flag.Parse()
    col := session.DB("jef_db").C("jef_collection")

    //Write MongoDB
    start_time := time.Now().Unix()
    for i := 0; i < *count; i++ {
        col.Insert(&User{
            ID:          bson.NewObjectId(),
            Username:    shuffle(usernames)[0],
            Age:         rand.Intn(81-13) + 13,
            Sex:         rand.Intn(2),
            Skill:       shuffle(skills)[:rand.Intn(6)],
            CreatedTime: time.Now(),
        })
    }

    end_time := time.Now().Unix() - start_time
    fmt.Printf("it took %v seconds writing datasets to mongodb...\n", end_time)
    fmt.Println()

    //Read MongoDB
    iter := col.Find(nil).Limit(10).Iter()
    for iter.Next(&tmp_user) {
        tmp_info = append(tmp_info, tmp_user)
    }
    json_data, err := json.MarshalIndent(tmp_info, "", "  ")
    if err != nil {
        log.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Println(string(json_data))
}
