package models

import "gopkg.in/mgo.v2/bson"

//Planet model
type Planet struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	Weather       string        `bson:"weather" json:"weather"`
	Ground        string        `bson:"ground" json:"ground"`
	AppearInFilms int           `bson:"appear_in_filmes" json:"appearInFilms"`
}

//FilterPlanet model
type FilterPlanet struct {
	Name *string
}
