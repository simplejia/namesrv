package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/simplejia/namesrv/mongo"
)

func (relation *Relation) GetsFromIp() (rels []*Relation, err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("index").C("relation")

	q := bson.M{
		"weight": bson.M{
			"$gt": 0,
		},
		"ip": relation.Ip,
	}

	sel := bson.M{
		"name":   1,
		"ip":     1,
		"port":   1,
		"udp":    1,
		"weight": 1,
	}
	err = c.Find(q).Select(sel).Sort("_id").All(&rels)
	if err != nil {
		return
	}

	return
}
