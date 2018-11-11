package model

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/simplejia/clog"
)

func (stat *Stat) CleanNumDay() (err error) {
	c := stat.GetC()
	defer c.Database.Session.Close()

	day := time.Now().Add(time.Hour * 24).Day()
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
