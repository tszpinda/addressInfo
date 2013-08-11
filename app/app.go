package main

import (
	"code.google.com/p/gorest"
	"fmt"
	address "github.com/tszpinda/addressInfo/address"
	distance "github.com/tszpinda/addressInfo/distance"
	keepAlive "github.com/tszpinda/addressInfo/keepAlive"
	view "github.com/tszpinda/addressInfo/web"
	"net/http"
	"os"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	ds := new(distance.DistanceService)
	gorest.RegisterService(ds)
	gorest.RegisterService(new(address.AddressService))

	serveSingle("/favicon.ico", "./favicon.ico")
	view.Mount()

	http.Handle("/", gorest.Handle())
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("starting app on port: " + port)
	keepAlive.Heroku(port)
	http.ListenAndServe(":"+port, nil)
}
