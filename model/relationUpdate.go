package model

func (relation *Relation) Update() (err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

	err = c.UpdateId(relation.Id, relation)
	if err != nil {
		return
	}
	return
}
