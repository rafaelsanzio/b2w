package services

import (
	"log"

	"../models"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

//PlanetsService to execute methods in databse
type PlanetsService struct {
	Server   string
	Database string
}

var db *mgo.Database

//Const to use in application
const (
	COLLECTION = "planets"
)

//Connect to make a connection with database
func (p *PlanetsService) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

//List to list all planets in database
func (p *PlanetsService) List(f models.FilterPlanet) ([]models.Planet, error) {
	var planets []models.Planet

	filter := bson.M{}
	if f.Name != nil {
		filter = bson.M{"name": bson.RegEx{*f.Name, "i"}}
	}

	err := db.C(COLLECTION).Find(filter).All(&planets)
	return planets, err
}

//GetByID to get a planet by id
func (p *PlanetsService) GetByID(id string) (*models.Planet, error) {
	var planet *models.Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

//Create to insert a planet in database
func (p *PlanetsService) Create(planet models.Planet) error {
	err := db.C(COLLECTION).Insert(&planet)
	return err
}

//Delete to delete a planet in database
func (p *PlanetsService) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

//Update to updated a planet in database
func (p *PlanetsService) Update(id string, planet models.Planet) error {
	planet.ID = bson.ObjectIdHex(id)

	err := db.C(COLLECTION).UpdateId(planet.ID, &planet)
	return err
}
