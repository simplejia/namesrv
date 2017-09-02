package controller

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/clog"
	"github.com/simplejia/namesrv/model"
)

// @postfilter("Boss")
func (relation *Relation) GetsFromIp(w http.ResponseWriter, r *http.Request) {
	fun := "relation.GetsFromIp"

	ip := r.FormValue("ip")
	cc := r.FormValue("cc")

	relationModel := model.NewRelation()
	relationModel.Ip = ip

	rels, err := relationModel.GetsFromIp()
	if err != nil {
		clog.Error("%s gets error: %v, req: %v", fun, err, relationModel)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ccNew := model.Relations(rels).CheckCode()
	if cc == ccNew {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"rels": rels,
		"cc":   ccNew,
	})
	return
}
