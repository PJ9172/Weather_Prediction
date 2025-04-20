package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type WeatherInput struct {
	Tmin float64 `json:"tmin"`
	Tmax float64 `json:"tmax"`
	Wspd float64 `json:"wspd"`
	Pres float64 `json:"pres"`
	Prcp float64 `json:"prcp"`
}

type PredictionResponse struct {
	PredictedTavg float64 `json:"predicted_tavg"`
}

func parseFloat(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func PredictWeather(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	input := WeatherInput{
		Tmin: parseFloat(r.FormValue("tmin")),
		Tmax: parseFloat(r.FormValue("tmax")),
		Wspd: parseFloat(r.FormValue("wspd")),
		Pres: parseFloat(r.FormValue("pres")),
		Prcp: parseFloat(r.FormValue("prcp")),
	}

	jsonData, _ := json.Marshal(input)

	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Prediction service error", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result PredictionResponse
	json.Unmarshal(body, &result)

	tmpl, _ := template.ParseFiles("templates/result.html","templates/layout.html")
	tmpl.Execute(w, result)
}
