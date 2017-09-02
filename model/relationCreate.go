package model

import "github.com/simplejia/namesrv/mongo"

func (relation *Relation) Create() (err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("index").C("relation")

	err = c.Insert(relation)
	if err != nil {
		return
	}
	return
}
