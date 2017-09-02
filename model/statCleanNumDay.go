package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/simplejia/clog"
	"github.com/simplejia/namesrv/mongo"
)

func (stat *Stat) CleanNumDay() (err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("stat").C("num_day")

	now := time.Now()
	day := now.Add(time.Hour * 24).Day()
	field := fmt.Sprintf("num_day_%d", day)
	sel := bson.M{
		field: bson.M{
			"$ne": 0,
		},
	}
	up := bson.M{
		"$set": bson.M{
			field: 0,
		},
	}
	_, err = c.UpdateAll(sel, up)
	if err != nil {
		return
	}

	return
}

func cleanNumDayTimer() {
	fun := "model.cleanNumDayTimer"

	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			err := NewStat().CleanNumDay()
			if err != nil {
				clog.Error("%s CleanNumDay err: %v", fun, err)
			}
		}
	}
}

func init() {
	go cleanNumDayTimer()
}
