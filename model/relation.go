package model

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type Relation struct {
	Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string
	Ip      string
	Port    int
	Udp     bool
	Weight  int
	Off     bool
	Creator string   `json:",omitempty"`
	Admins  []string `json:",omitempty"`
	Ct      int64    `json:",omitempty"`
	Ut      int64    `json:",omitempty"`
}

func (relation *Relation) Regular() (ok bool) {
	if relation == nil {
		return
	}

	relation.Name = strings.TrimSpace(relation.Name)
	if relation.Name == "" {
		return
	}

	relation.Ip = strings.TrimSpace(relation.Ip)
	if relation.Ip == "" {
		return
	}

	if relation.Port <= 0 {
		return
	}

	relation.Creator = strings.TrimSpace(relation.Creator)
	if relation.Creator == "" {
		return
	}

	ok = true
	return
}

func NewRelation() *Relation {
	rel := &Relation{}
	return rel
}

type Relations []*Relation

func (relations Relations) CheckCode() (cc string) {
	if len(relations) == 0 {
		return
	}

	ctx := md5.New()
	for _, relation := range relations {
		io.WriteString(ctx, fmt.Sprintf(
			"%s,%d,%t,%d,%t",
			relation.Ip,
			relation.Port,
			relation.Udp,
			relation.Weight,
			relation.Off,
		))
	}
	cc = fmt.Sprintf("%x", ctx.Sum(nil))

	return
}
