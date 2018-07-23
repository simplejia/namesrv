package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/simplejia/clog"
	"github.com/simplejia/namesrv/model"
)

// @postfilter("Boss")
func (relation *Relation) Update(w http.ResponseWriter, r *http.Request) {
	fun := "relation.Update"

	body := relation.ReadBody(r)

	relationModel := model.NewRelation()
	err := json.Unmarshal(body, relationModel)
	if err != nil || !relationModel.Regular() {
		clog.Error("%s unmarshal error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	relationModel.Ut = time.Now().Unix()
	err = relationModel.Update()
	if err != nil {
		clog.Error("%s update error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	relation.ReplyOk(w, relationModel)
	return
}
