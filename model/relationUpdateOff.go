package model

import (
	"gopkg.in/mgo.v2/bson"
)

func (relation *Relation) UpdateOff() (err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

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
