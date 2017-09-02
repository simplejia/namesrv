package model

type Stat struct {
	Name   string
	NumDay int `json:"num_day,omitempty"`
}

func (stat *Stat) Regular() (ok bool) {
	if stat == nil {
		return
	}

	ok = true
	return
}

func NewStat() *Stat {
	stat := &Stat{}
	return stat
}
