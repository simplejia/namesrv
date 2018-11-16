package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/namesrv/model"
)

// @postfilter("Boss")
func (relation *Relation) Create(w http.ResponseWriter, r *http.Request) {
	fun := "relation.Create"

	body := relation.ReadBody(r)

	relationModel := model.NewRelation()
	err := json.Unmarshal(body, relationModel)
	if err != nil || !relationModel.Regular() {
		clog.Error("%s unmarshal error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	relationModel.Ct = time.Now().Unix()
	err = relationModel.Create()
	if err != nil {
		clog.Error("%s create error: %v, req: %s", fun, err, body)
		relation.ReplyFail(w)
		return
	}

	relation.ReplyOk(w, relationModel)
	return
}
