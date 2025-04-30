package main

import (
	"log"
	"net/http"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/api"
)

func main()  {
	config.Init() //環境変数の読み込み

	router := api.NewRouter()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

