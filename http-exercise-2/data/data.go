package data

import (
        "errors"
        "github.com/google/uuid"
        "golang.org/x/crypto/bcrypt"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "time"
)

var (
        session      = &mgo.Session{}
        col_user     = &mgo.Collection{}
        col_session  = &mgo.Collection{}
        mgo_server1  = "10.5.1.160:27017"
        mgo_database = "admin"
        mgo_username = "jef"
        mgo_password = "mongopw"
        mgo_timeout  = 6 //seconds
)

type User struct {
        Id        bson.ObjectId `bson:"_id,omitempty"`
        Uuid      string        `bson:"uuid"`
        Name      string        `bson:"name"`
        Password  string        `bson:"password"`
        Sex       string        `bson:"sex"`
        Age       int           `bson:"age"`
        Email     []string      `bson:"email"`
        CreatedAt time.Time     `bson:"created_at"`
}

type Session struct {
        Id        bson.ObjectId `bson:"_id,omitempty"`
        Uuid      string        `bson:"uuid"`
        UserName  string        `bson:"user_name"`
        UserId    string        `bson:"user_id"`
        CreatedAt time.Time     `bson:"created_at"`
}

func (u *User) Create() (err error) {
        count, err := col_user.Find(bson.M{"name": u.Name}).Count()

        if err != nil {
                return
        } else if count >= 1 {
                return errors.New("data exist")
        } else if count == 0 {
                u.Id = bson.NewObjectId()
                u.Uuid = uuid.New().String()

                hashed_password, err1 := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
                if err1 != nil {
                        return
                }
                u.Password = string(hashed_password)
                u.CreatedAt = time.Now()

                col_user.Insert(u)
        }
        return
}

func GetUserByName(user_name string) (u *User, err error) {
        err = col_user.Find(bson.M{"name": user_name}).One(&u)
        if err != nil {
                return
        }
        return
}

func GetUserByUuid(user_uuid string) (u *User, err error) {
        err = col_user.Find(bson.M{"uuid": user_uuid}).One(&u)
        if err != nil {
                return
        }
        return
}

func GetUserBySessionUuid(session_uuid string) (u *User, err error) {
        temp_session := Session{}

        count, err := col_session.Find(bson.M{"uuid": session_uuid}).Count()
        if err != nil {
                return
        } else if count >= 1 {
                err = col_session.Find(bson.M{"uuid": session_uuid}).One(&temp_session)
                if err != nil {
                        return
                }
        } else if count == 0 {
                err = errors.New("not found")
                return
        }

        err = col_user.Find(bson.M{"uuid": temp_session.UserId}).One(&u)
        if err != nil {
                return
        }
        return
}

func (u *User) Update() (err error) {
        count, err := col_user.Find(bson.M{"uuid": u.Uuid}).Count()
        if err != nil {
                return
        } else if count >= 1 {
                //update info goes here
                store_user := &User{}
                err1 := col_user.Find(bson.M{"uuid": u.Uuid}).One(store_user)
                if err1 != nil {
                        return
                }
                if u.Password != "" {
                        store_user.Password = u.Password
                }
                if u.Sex != "" {
                        store_user.Sex = u.Sex
                }
                if u.Age != 0 {
                        store_user.Age = u.Age
                }
                if u.Email != nil {
                        store_user.Email = u.Email
                }

                col_user.Update(
                        bson.M{"uuid": u.Uuid},
                        bson.M{"$set": bson.M{
                                "password": store_user.Password,
                                "sex":      store_user.Sex,
                                "age":      store_user.Age,
                                "email":    store_user.Email,
                        }})
        } else if count == 0 {
                return errors.New("not found")
        }
        return
}

func (u *User) Delete() (err error) {
        count, err := col_user.Find(bson.M{"uuid": u.Uuid}).Count()
        if err != nil {
                return
        } else if count >= 1 {
                //delete code goes here
                err = col_user.Remove(bson.M{"uuid": u.Uuid})
        } else if count == 0 {
                return errors.New("not fount")
        }
        return
}

func (u *User) CreateSession() (s *Session, err error) {
        count, err := col_session.Find(bson.M{"user_id": u.Uuid}).Count()
        if err != nil {
                return
        } else if count >= 1 {
                err = errors.New("data exist")
                return
        } else if count == 0 {
                s = &Session{
                        Id:        bson.NewObjectId(),
                        Uuid:      uuid.New().String(),
                        UserName:  u.Name,
                        UserId:    u.Uuid,
                        CreatedAt: time.Now(),
                }

                err = col_session.Insert(s)
                if err != nil {
                        return
                }
                return
        }
        return
}

func (u *User) Session() (s *Session, err error) {
        count, err := col_session.Find(bson.M{"user_id": u.Uuid}).Count()
        if err != nil {
                return
        } else if count >= 1 {
                err = col_session.Find(bson.M{"user_id": u.Uuid}).One(&s)
                if err != nil {
                        return
                }
        } else if count == 0 {
                err = errors.New("not found")
                return
        }
        return
}

func (s *Session) User() (u *User, err error) {
        temp_session := Session{}
        count, err := col_session.Find(bson.M{"uuid": s.Uuid}).Count()
        if err != nil {
                return
        } else if count == 0 {
                err = errors.New("not found")
        } else if count >= 1 {
                err = col_session.Find(bson.M{"uuid": s.Uuid}).One(&temp_session)
                if err != nil {
                        return
                }
                err = col_user.Find(bson.M{"uuid": temp_session.UserId}).One(&u)
                if err != nil {
                        return
                }
        }
        return
}

func (s *Session) Delete() (err error) {
        count, err := col_session.Find(bson.M{"uuid": s.Uuid}).Count()
        if err != nil {
                return
        } else if count == 0 {
                return errors.New("not found")
        } else if count >= 1 {
                err = col_session.Remove(bson.M{"uuid": s.Uuid})
                if err != nil {
                        return
                }
        }
        return
}

func CloseMgoSession() {
        session.Close()
}

func init() {
        session, err := mgo.DialWithInfo(&mgo.DialInfo{
                Addrs:    []string{mgo_server1},
                Database: mgo_database,
                Username: mgo_username,
                Password: mgo_password,
                Timeout:  time.Duration(mgo_timeout) * time.Second,
        })
        if err != nil {
                panic(err)
        }
        col_user = session.DB("jef_db").C("user")
        col_session = session.DB("jef_db").C("session")
}

