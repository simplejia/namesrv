package model

func (relation *Relation) Create() (err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(relation)
	if err != nil {
		return
	}
	return
}
