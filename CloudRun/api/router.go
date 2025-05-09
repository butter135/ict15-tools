package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/middleware"
)

func NewRouter() http.Handler{
	r := mux.NewRouter()
	r.Use(middleware.NormalizeInput) // ← ここで先に成型！

	r.HandleFunc("/dice", HandleDiceRoll).Methods("POST")
	r.HandleFunc("/picchan", HandlePicChan).Methods("POST")
	r.HandleFunc("/tokumei", HandleTokumei).Methods("POST")
	r.HandleFunc("/twifix", HandleTwiFix).Methods("POST")
	//r.HandleFuncで追加してく
	return r
}