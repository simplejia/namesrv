package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (stat *Stat) IncNumDay(num int) (err error) {
	c := stat.GetC()
	defer c.Database.Session.Close()

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
