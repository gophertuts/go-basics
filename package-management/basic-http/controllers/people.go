package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func people(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "Hello Mr. %s!\n", ps.ByName(nameParam))
	if err != nil {
		log.Fatal(err)
	}
}
