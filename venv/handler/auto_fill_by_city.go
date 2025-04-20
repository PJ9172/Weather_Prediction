package handler

import (
	"encoding/json"
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

func AutoFillByCity(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		city := r.FormValue("city")
		apiid := os.Getenv("API_ID")
		url := "https://api.openweathermap.org/data/2.5/weather?appid=" + apiid + "&units=metric&q=" + city

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "City not found", http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		var weather Input
		json.Unmarshal(body, &weather)

		// fmt.Println(string(body))

		data := map[string]float64{
			"tmin": weather.Main.TempMin,
			"tmax": weather.Main.TempMax,
			"wspd": weather.Wind.Speed,
			"pres": weather.Main.Pressure,
			"prcp": weather.Rain.OneH,
		}

		tmpl, err := template.ParseFiles("templates/autofill.html", "templates/layout.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	}
}
