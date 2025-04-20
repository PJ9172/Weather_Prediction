package router

import (
	"Weather_Prediction/venv/handler"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() *mux.Router {
	r := mux.NewRouter() //create router object
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/index.html", "templates/layout.html")
		tmpl.Execute(w, nil)
	}) //handel root url

	r.HandleFunc("/autofill", handler.AutoFillByCity).Methods("POST")
	r.HandleFunc("/predict", handler.PredictWeather).Methods("POST")

	return r
}
