package model

func (relation *Relation) Count() (n int, err error) {
	c := relation.GetC()
	defer c.Database.Session.Close()

	n, err = c.Count()
	if err != nil {
		return
	}
	return
}
