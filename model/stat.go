package model

import (
	mgo "github.com/globalsign/mgo"
	"github.com/simplejia/namesrv/mongo"
)

type Stat struct {
	Name string
}

func (stat *Stat) Regular() (ok bool) {
	if stat == nil {
		return
	}

	ok = true
	return
}

// Db 返回db name
func (stat *Stat) Db() (db string) {
	return "stat"
}

// Table 返回table name
func (stat *Stat) Table() (table string) {
	return "num_day"
}

// GetC 返回db col
func (stat *Stat) GetC() (c *mgo.Collection) {
	db, table := stat.Db(), stat.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}

func NewStat() *Stat {
	stat := &Stat{}
	return stat
}
