package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/simplejia/namesrv/mongo"
)

func (relation *Relation) UpdateOff() (err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("index").C("relation")

	sel := bson.M{
		"ip":   relation.Ip,
		"port": relation.Port,
	}
	up := bson.M{
		"$set": bson.M{
			"off": relation.Off,
		},
	}
	_, err = c.UpdateAll(sel, up)
	if err != nil {
		return
	}

	return
}
