package view

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func Mount() {
	pfx := "/public/static/img/"
	h := http.StripPrefix(pfx, http.FileServer(http.Dir("./public/static/img")))
	http.Handle(pfx, h)
	
	serveSingle("/static/js/main.js", "./public/static/js/main.js")
 	serveSingle("/static/js/jquery.js", "./public/static/js/jquery.js")
 	serveSingle("/static/bootstrap/js/bootstrap.js", "./public/static/bootstrap/js/bootstrap.js")
 	serveSingle("/static/bootstrap/css/bootstrap.css", "./public/static/bootstrap/css/bootstrap.min.css")
 	serveSingle("/static/bootstrap/css/bootstrap-responsive.css", "./public/static/bootstrap/css/bootstrap-responsive.min.css")
 	serveSingle("/static/bootstrap/css/main.css", "./public/static/bootstrap/css/main.css")

	http.HandleFunc("/index", index)
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templateName = "public/tmpl/search.gtpl"
		t, err := template.ParseFiles(templateName)
		if err != nil {
			log.Fatal("unable to parse", templateName, err)
		}
		t.Execute(w, nil)
	}
}
