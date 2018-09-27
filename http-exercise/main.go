package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
)

var (
    mongo_db  = "jef_db"
    mongo_col = "jef_col"
    session   = &mgo.Session{}
    col       = &mgo.Collection{}
    us        = UserStore{}
)

type User struct {
    Id          bson.ObjectId `bson:"_id,omitempty"`
    Name        string        `bson:"name"`
    Sex         int           `bson:"sex"`
    Age         int           `bson:"age"`
    Email       string        `bson:"email"`
    Ocupation   []string      `bson:"ocupation"`
    Title       string        `bson:"title"`
    Company     string        `bson:"company"`
    Tags        []string      `bson:"tags"`
    CreatedTime time.Time     `bson:"created_time"`
}

type UserStore struct {
    C *mgo.Collection
}

func (store UserStore) Create(user *User) error {
    var err error
    user.Id = bson.NewObjectId()
    user.CreatedTime = time.Now()

    if n, err := store.C.Find(bson.M{"name": user.Name}).Count(); err != nil {
        log.Printf("Error: %v", err)
    } else if n == 0 {
        err := store.C.Insert(*user)
        checkErr(err)
    } else if n >= 1 {
        log.Printf("Error: %v", errors.New("User existed! can not create..."))
    }
    return err
}

func (store UserStore) GetAll() []User {
    var user_temp User
    var users []User
    iter := store.C.Find(nil).Iter()

    for iter.Next(&user_temp) {
        users = append(users, user_temp)
    }
    return users
}

func (store UserStore) Update(user *User) error {
    err := store.C.Update(
        bson.M{"name": user.Name},
        bson.M{"$set": bson.M{
            "name":         user.Name,
            "sex":          user.Sex,
            "age":          user.Age,
            "email":        user.Email,
            "ocupation":    user.Ocupation,
            "title":        user.Title,
            "company":      user.Company,
            "tags":         user.Tags,
            "created_time": user.CreatedTime,
        }})
    return err
}

func (store UserStore) Delete(user *User) error {
    err := store.C.Remove(bson.M{"name": user.Name})
    return err
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func init() {
    session, err := mgo.DialWithInfo(&mgo.DialInfo{
        Addrs:    []string{"10.5.1.160"},
        Database: "admin",
        Username: "jef",
        Password: "mongopw",
        Timeout:  time.Duration(3) * time.Second,
    })
    checkErr(err)

    us.C = session.DB(mongo_db).C(mongo_col)
}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/register", 302)
}

func register(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("register.html")
    checkErr(err)
    t.Execute(w, "")
}

func postRegister(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    n, err := us.C.Find(bson.M{"name": r.PostFormValue("name")}).Count()
    if err != nil {
        fmt.Fprintf(w, "shit happend. Error: %v", err)
        log.Printf("Error: ", err)
    } else if n == 0 {
        sex, _ := strconv.Atoi(r.PostFormValue("sex"))
        age, _ := strconv.Atoi(r.PostFormValue("age"))

        err := us.Create(&User{
            Id:          bson.NewObjectId(),
            Name:        r.PostFormValue("name"),
            Sex:         sex, // int
            Age:         age, //int
            Email:       r.PostFormValue("email"),
            Ocupation:   []string{r.PostFormValue("ocupation")}, // []string
            Title:       r.PostFormValue("title"),
            Company:     r.PostFormValue("company"),
            Tags:        []string{r.PostFormValue("tags")}, // []string
            CreatedTime: time.Now(),
        })
        checkErr(err)
        w.Write([]byte("account successfully created."))
    } else if n >= 1 {
        w.Write([]byte("account has existed."))
    }
}
func getall(w http.ResponseWriter, r *http.Request) {
    var users []User
    users = us.GetAll()
    json_users, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(json_users))
}

func main() {
    defer session.Close()

    server := &http.Server{
        Addr: ":80",
    }

    http.HandleFunc("/", index)
    http.HandleFunc("/register", register)
    http.HandleFunc("/postRegister", postRegister)

    http.HandleFunc("/getall", getall)

    fmt.Println("Serving web app...")
    server.ListenAndServe()

}
