package controllers

import (
	"github.com/julienschmidt/httprouter"
)

const (
	nameParam = "name"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/people/:name", people)
	return router
}
