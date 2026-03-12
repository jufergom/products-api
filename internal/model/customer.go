package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Customer struct {
	ID     bson.ObjectID      `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Active bool               `bson:"active" json:"active"`
}
