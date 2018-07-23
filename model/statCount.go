package model

func (stat *Stat) Count() (n int, err error) {
	c := stat.GetC()
	defer c.Database.Session.Close()

	n, err = c.Count()
	if err != nil {
		return
	}
	return
}
