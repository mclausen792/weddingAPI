package dao

import (
	"log"

	. "weddingAPI/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GuestsDAO struct {
	Server   string
	Database string
	DialInfo *mgo.DialInfo
}

var db *mgo.Database

const (
	GUESTCOLLECTION = "Guests"
)

func (m *GuestsDAO) Connect() {
	session, err := mgo.DialWithInfo(m.DialInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
func (m *GuestsDAO) FindAllGuests() ([]Guests, error) {
	var guests []Guests
	err := db.C(GUESTCOLLECTION).Find(bson.M{}).All(&guests)
	return guests, err
}

func (m *GuestsDAO) Insert(guest Guests) error {
	err := db.C(GUESTCOLLECTION).Insert(&guest)
	return err
}
