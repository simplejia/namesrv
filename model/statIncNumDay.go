package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/simplejia/namesrv/mongo"
)

func (stat *Stat) IncNumDay(num int) (err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("stat").C("num_day")

	sel := bson.M{
		"name": stat.Name,
	}

	now := time.Now()
	day := now.Day()
	field := fmt.Sprintf("num_day_%d", day)
	up := bson.M{
		"$inc": bson.M{
			field: num,
		},
	}
	_, err = c.Upsert(sel, up)
	if err != nil {
		return
	}

	return
}
