package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		// handle errors bastard
		_, _ = fmt.Fprintf(w, "Welcome!")
	})
	router.GET("/people/:name", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		name := ps.ByName("name")
		// handle errors bastard
		_, _ = fmt.Fprintf(w, "Welcome %s", name)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
