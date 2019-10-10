package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprint(w, "Welcome to index!\n")
	if err != nil {
		log.Fatal(err)
	}
}
