package main

import (
    "encoding/json"
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type User struct {
    ID          bson.ObjectId `bson:"_id,omitempty"`
    Username    string        `json:"username"`
    Age         int           `json:"age"`
    Sex         string        `json:"sex"`
    Createdtime time.Time     `json:"created_time"`
}

type UserAdmin struct {
    C *mgo.Collection
}

func (m UserAdmin) Create(b *User) error {
    b.ID = bson.NewObjectId()
    err := m.C.Insert(b)
    return err
}

func (m UserAdmin) Read() error {
    var b []User
    t := User{}
    iter := m.C.Find(nil).Iter()
    for iter.Next(&t) {
        b = append(b, t)
    }
    json_data, err := json.MarshalIndent(b, "", "  ")
    if err != nil {
        return err
    } else {
        fmt.Println(string(json_data))
    }
    return nil
}

func (m UserAdmin) Update(b *User) error {
    err := m.C.Update(bson.M{"username": "jef"},
        bson.M{"$set": bson.M{
            "username": b.Username,
            "age":      b.Age,
            "sex":      b.Sex,
        }})
    return err
}

func (m UserAdmin) Delete() error {
    _, err := m.C.RemoveAll(nil)
    return err
}

func main() {
    mongoDialInfo := &mgo.DialInfo{
        Addrs:    []string{"10.5.1.160"},
        Timeout:  60 * time.Second,
        Username: "jef",
        Password: "mongopw",
        Database: "admin",
    }

    session, _ := mgo.DialWithInfo(mongoDialInfo)
    user_col := session.DB("jef_db").C("jef_collection")

    ua := UserAdmin{C: user_col}
    ua.Create(&User{Username: "jef", Age: 38, Sex: "M", Createdtime: time.Now()})
    ua.Update(&User{Age: 8, Sex: "F", Username: "shit"})
    ua.Read()
    ua.Delete()
}
