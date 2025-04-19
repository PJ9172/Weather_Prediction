package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Input struct {
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
}

func FetchData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		city := r.FormValue("city")
		// fmt.Println(city)
		apikey := os.Getenv("API_KEY")
		apiurl := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apikey + "&units=metric"

		resp, err := http.Get(apiurl) //API call
		if err != nil {
			http.Error(w, "City Not Found!!!", resp.StatusCode)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body) //Read Data
		var input Input
		json.Unmarshal(body, &input) //Convert only needed JSON data to Struct
		fmt.Println(input)

		inputJson, err := json.MarshalIndent(input, " ", "\t") //Convert Struct to JSON
		if err != nil {
			http.Error(w, "Error to convert Struct to Json!!!", http.StatusInternalServerError)
			return
		}
		resp, err = http.Post("http://localhost:3000/predict", "application/json", bytes.NewBuffer(inputJson))
		if err != nil {
			http.Error(w, "Failed in Post request to FastAPI!!!", http.StatusInternalServerError)
			return
		}
		// defer resp.Body.Close()

		body, _ = ioutil.ReadAll(resp.Body)

		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, "Template Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, body)
	}
}

func Prediction(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/predict_form.html")
	tmpl.Execute(w, nil)

}
