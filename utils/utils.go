package utils

import (
	"encoding/json"
	"net/http"

	"../models"
)

//GetAppereanceInFilms getting planets from api starwars
func GetAppereanceInFilms(name string) (int, error) {
	response := models.ResponseSwapi{}
	number := 0
	resp, err := http.Get("https://swapi.dev/api/planets/?search=" + name)
	if err != nil {
		return number, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return number, err
	}

	if len(response.Result) > 0 {
		number = len(response.Result[0].Films)
	}

	return number, nil
}

//ResponseWithError response API with error
func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	ResponseWithJSON(w, code, map[string]string{"error": msg})
}

//ResponseWithJSON response API with JSON
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
