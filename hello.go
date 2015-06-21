package hello

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	guestbookForm []byte
	signTemplate  = template.Must(template.ParseFiles("guestbook.html"))
)

func init() {
	content, err := ioutil.ReadFile("guestbookform.html")
	if err != nil {
		panic(err)
	}
	guestbookForm = content

	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) { w.Write(guestbookForm) }

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// package hello
//
// import (
// 	"fmt"
// 	"net/http"
//
// 	"appengine"
// 	"appengine/user"
// )
//
// func init() {
// 	http.HandleFunc("/", handler)
// }
//
// func handler(w http.ResponseWriter, r *http.Request) {
//
// 	c := appengine.NewContext(r)
// 	u := user.Current(c)
// 	if u == nil {
// 		url, err := user.LoginURL(c, r.URL.String())
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Location", url)
// 		w.WriteHeader(http.StatusFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello, %v! You look messed up today :p", u)
//
// 	//fmt.Fprint(w, "Hello, GoAppEngine!")
// }
