package model

func (stat *Stat) List(offset, limit int) (l []map[string]interface{}, err error) {
	c := stat.GetC()
	defer c.Database.Session.Close()

	err = c.Find(nil).Sort("name").Skip(offset).Limit(limit).All(&l)
	if err != nil {
		return
	}
	return
}
