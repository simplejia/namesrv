package model

func (relation *Relation) List(offset, limit int) (relations []*Relation, err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

	err = c.Find(nil).Sort("name").Skip(offset).Limit(limit).All(&relations)
	if err != nil {
		return
	}
	return
}
