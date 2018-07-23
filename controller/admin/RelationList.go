package admin

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/clog"
	"github.com/simplejia/namesrv/model"
)

type RelationListReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// Regular 用于参数校验
func (relationListReq *RelationListReq) Regular() (ok bool) {
	if relationListReq == nil {
		return
	}

	if relationListReq.Limit <= 0 {
		relationListReq.Limit = 20
	}

	ok = true
	return
}

// @postfilter("Boss")
func (relation *Relation) List(w http.ResponseWriter, r *http.Request) {
	fun := "relation.List"

	var relationListReq *RelationListReq
	if err := json.Unmarshal(relation.ReadBody(r), &relationListReq); err != nil || !relationListReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, relationListReq)
		relation.ReplyFail(w)
		return
	}

	n, err := model.NewRelation().Count()
	if err != nil {
		clog.Error("%s count error: %v, req: %v", fun, err, relationListReq)
		relation.ReplyFail(w)
		return
	}

	relations, err := model.NewRelation().List(relationListReq.Offset, relationListReq.Limit)
	if err != nil {
		clog.Error("%s list error: %v, req: %v", fun, err, relationListReq)
		relation.ReplyFail(w)
		return
	}

	relation.ReplyOk(w, map[string]interface{}{
		"list":   relations,
		"total":  n,
		"offset": relationListReq.Offset + len(relations),
	})
	return
}
