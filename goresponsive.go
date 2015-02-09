/*
* Go Responsive - boilerplate for Go programming language
* Based on http://www.alexedwards.net/blog/serving-static-sites-with-go
 */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	lp := path.Join("static", "index.html")
	tmpl, err := template.ParseFiles(lp)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	if err := tmpl.ExecuteTemplate(w, "home", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/home/", serveTemplate)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
