package admin

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/namesrv/model"
)

// @postfilter("Boss")
func (relation *Relation) Delete(w http.ResponseWriter, r *http.Request) {
	fun := "relation.Delete"

	body := relation.ReadBody(r)

	relationModel := model.NewRelation()
	err := json.Unmarshal(body, relationModel)
	if err != nil {
		clog.Error("%s unmarshal error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	err = relationModel.Delete()
	if err != nil {
		clog.Error("%s create error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	relation.ReplyOk(w, nil)
	return
}
