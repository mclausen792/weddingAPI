package models

import "gopkg.in/mgo.v2/bson"

type Guests struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	First     string        `bson:"first" json:"first"`
	Last      string        `bson:"last" json:"last"`
	Attending string        `bson:"attending" json:"attending"`
	Children  int           `bson:"children" json: "children"`
	Adults    int           `bson:"adults" json:"adults"`
}
