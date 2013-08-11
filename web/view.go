package view

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func Mount() {
	pfx := "/static/img/"
	h := http.StripPrefix(pfx, http.FileServer(http.Dir("./static/img")))
	http.Handle(pfx, h)
	
	serveSingle("/static/js/main.js", "./static/js/main.js")
 	serveSingle("/static/js/jquery.js", "./static/js/jquery.js")
 	serveSingle("/static/bootstrap/js/bootstrap.js", "./static/bootstrap/js/bootstrap.js")
 	serveSingle("/static/bootstrap/css/bootstrap.css", "./static/bootstrap/css/bootstrap.min.css")
 	serveSingle("/static/bootstrap/css/bootstrap-responsive.css", "./static/bootstrap/css/bootstrap-responsive.min.css")
	http.HandleFunc("/index", index)
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templateName = "tmpl/search.gtpl"
		t, err := template.ParseFiles(templateName)
		if err != nil {
			log.Fatal("unable to parse", templateName, err)
		}
		t.Execute(w, nil)
	}
}
