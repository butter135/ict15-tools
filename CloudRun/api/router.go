package api

import (
	"net/http"
	"github.com/gorilla/mux"

)

func NewRouter() http.Handler{
	r := mux.NewRouter()

	r.HandleFunc("/dice", HandleDiceRoll).Methods("POST")

	//r.HandleFuncで追加してく
	return r
}