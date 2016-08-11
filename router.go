package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	api := router.PathPrefix("/api").Subrouter();
	api.HandleFunc("/menu/{type}", MenuHandler)
	api.HandleFunc("/sort/{type}/{property}", SortHandler)
	api.HandleFunc("/value/{type}/{property}", ValueHandler)
	api.HandleFunc("/value/{money}/{type}/{property}", MoneyValueHandler)
	return router
}
