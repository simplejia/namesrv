package model

func (relation *Relation) Delete() (err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

	err = c.RemoveId(relation.Id)
	if err != nil {
		return
	}

	return
}
