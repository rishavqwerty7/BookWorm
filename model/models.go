package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bookworm struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"id,omitempty"`
	Name   string             `json:"name,omitempty"`
	Author string             `json:"author,omitempty"`
	Read   bool               `json:"read,omitempty"`
}
