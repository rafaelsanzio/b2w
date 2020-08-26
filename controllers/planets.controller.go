package controllers

import (
	"encoding/json"
	"net/http"

	conf "../config"
	"../models"
	"../services"
	"../utils"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

var planetServices = services.PlanetsService{}
var config = conf.Config{}

func init() {
	config.Read()

	planetServices.Server = config.Server
	planetServices.Database = config.Database
	planetServices.Connect()
}

//List function to list all planets
func List(w http.ResponseWriter, r *http.Request) {
	f := models.FilterPlanet{}
	params := r.URL.Query()

	if len(params["name"]) > 0 {
		name := params["name"][0]
		f.Name = &name
	}
	planets, err := planetServices.List(f)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, planets)
}

//GetByID function to get a planet by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := planetServices.GetByID(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid Planet ID")
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, planet)
}

//Create function to insert a planet in database
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet models.Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	appearInFilms, err := utils.GetAppereanceInFilms(planet.Name)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	planet.ID = bson.NewObjectId()
	planet.AppearInFilms = appearInFilms

	if err := planetServices.Create(planet); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusCreated, planet)
}

//Update function to update a planet in database
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet models.Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	appearInFilms, err := utils.GetAppereanceInFilms(planet.Name)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	planet.AppearInFilms = appearInFilms

	if err := planetServices.Update(params["id"], planet); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, map[string]string{"result": planet.Name + " atualizado com sucesso!"})
}

//Delete dunction to remove a planet in database
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if err := planetServices.Delete(params["id"]); err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
