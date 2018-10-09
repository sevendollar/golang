package main

import (
        "fmt"
        "golang.org/x/crypto/bcrypt"
        "html/template"
        "jef/data"
        "net/http"
        "strconv"
        "time"
)

func index(w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles("templates/index.html")
        if err != nil {
                return
        }
        t.Execute(w, "")
}

func signup(w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles("templates/signup.html")
        if err != nil {
                return
        }
        t.Execute(w, "")
}

func process_signup(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()

        age, err := strconv.Atoi(r.PostFormValue("age"))
        if err != nil {
                return
        }

        u := &data.User{
                Name:     r.PostFormValue("name"),
                Password: r.PostFormValue("password"),
                Sex:      r.PostFormValue("sex"),
                Age:      age,
                Email:    []string{r.PostFormValue("email")},
        }
        err = u.Create()
        if err != nil {
                fmt.Fprintf(w, "failed: %v", err)
                return
        } else {
                fmt.Fprintf(w, "suceed")
        }

}

func signin(w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles("templates/signin.html")
        if err != nil {
                return
        }
        t.Execute(w, "")
}

func signout(w http.ResponseWriter, r *http.Request) {
        session_id, err := r.Cookie("SessionId")
        if err != nil {
                fmt.Fprintln(w, "you're already sign-outed!")
                return
        }

        s := &data.Session{
                Uuid: session_id.Value,
        }
        err = s.Delete()
        if err != nil {
                return
        }

        http.SetCookie(w, &http.Cookie{
                Name:    "SessionId",
                MaxAge:  -1,
                Expires: time.Unix(1, 0),
        })
        fmt.Fprintln(w, "you're sign-outed!")
}

func autheticate(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()

        post_name := r.PostFormValue("name")
        post_password := r.PostFormValue("password")
        store_user, err := data.GetUserByName(post_name)
        if err != nil {
                fmt.Fprintln(w, "Error:", err)
                return
        }
        hashed_passwod := store_user.Password

        if err := bcrypt.CompareHashAndPassword([]byte(hashed_passwod), []byte(post_password)); err != nil {
                //failed code goes here
                fmt.Fprintln(w, "Sign In Failed!, Error:", err)
        } else {
                //sucseed code goes here

                session, err := store_user.CreateSession()
                if err != nil {
                        fmt.Fprintln(w, "Error:", err)
                        return
                }

                http.SetCookie(w, &http.Cookie{
                        Name:     "SessionId",
                        Value:    session.Uuid,
                        HttpOnly: true,
                })

                fmt.Fprintln(w, "Sign In Succeed!")
        }
}

func verify(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()

        cookie_map := getCookie(r.Header.Get("cookie"))
        session_id := cookie_map["SessionId"]

        u, err := data.GetUserBySessionUuid(session_id)
        if err != nil {
                fmt.Fprintln(w, "hi guest")
                return
        }
        fmt.Fprintf(w, "hi %v, you're now %v years old.", u.Name, u.Age)
}

