package model

import "github.com/simplejia/namesrv/mongo"

func (relation *Relation) Delete() (err error) {
	session := mongo.DBS["index"]
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB("index").C("relation")

	err = c.RemoveId(relation.Id)
	if err != nil {
		return
	}

	return
}
