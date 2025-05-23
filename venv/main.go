package main

import (
	"Weather_Prediction/venv/config"
	"Weather_Prediction/venv/router"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	r := router.SetRouter() //routing function

	log.Println("Server running on 3000 port!!!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
