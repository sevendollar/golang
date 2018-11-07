package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"web/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("SessionId")
	if err != nil {
		t, err := template.ParseFiles("/go/src/web/templates/index_guest.html")
		if err != nil {
			return
		}
		t.Execute(w, "")
	} else {
		u := &data.User{}
		u, err = data.GetUserBySessionUuid(session_id.Value)
		if err != nil {
			t, err := template.ParseFiles("/go/src/web/templates/index_guest.html")
			if err != nil {
				return
			}
			t.Execute(w, "")
			return
		}
		t, err := template.ParseFiles("/go/src/web/templates/index.html")
		if err != nil {
			return
		}
		t.Execute(w, u)

		t, err = template.ParseFiles("/go/src/web/templates/getAllUsers.html")
		if err != nil {
			return
		}
		users := []data.User{}
		users, err = data.GetAllUsers()
		if err != nil {
			return
		}
		t.Execute(w, users)

	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/go/src/web/templates/signup.html")
	if err != nil {
		return
	}
	t.Execute(w, "")
}

func process_signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	age, err := strconv.Atoi(r.PostFormValue("age"))
	if err != nil {
		fmt.Fprintf(w, "something went wrong with Age field, Error: %v", err)
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
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		fmt.Fprintf(w, "suceed")
	}

}

func signin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/go/src/web/templates/signin.html")
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
	http.Redirect(w, r, "/", 302)
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

		http.Redirect(w, r, "/", 302)
	}
}

func verify(w http.ResponseWriter, r *http.Request) {
	cookie_map := getCookie(r.Header.Get("cookie"))
	session_id := cookie_map["SessionId"]

	u, err := data.GetUserBySessionUuid(session_id)
	if err != nil {
		fmt.Fprintln(w, "hi guest")
		return
	}
	fmt.Fprintf(w, "hi %v, you're now %v years old and your UUID is: %v", u.Name, u.Age, u.Uuid)
}
