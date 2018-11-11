package model

import (
	"github.com/globalsign/mgo/bson"
)

func (relation *Relation) GetsFromIp() (rels []*Relation, err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

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
