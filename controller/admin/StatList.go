package admin

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/namesrv/model"
)

type StatListReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// Regular 用于参数校验
func (statListReq *StatListReq) Regular() (ok bool) {
	if statListReq == nil {
		return
	}

	if statListReq.Limit <= 0 {
		statListReq.Limit = 20
	}

	ok = true
	return
}

// @postfilter("Boss")
func (stat *Stat) List(w http.ResponseWriter, r *http.Request) {
	fun := "stat.List"

	var statListReq *StatListReq
	if err := json.Unmarshal(stat.ReadBody(r), &statListReq); err != nil || !statListReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, statListReq)
		stat.ReplyFail(w)
		return
	}

	n, err := model.NewStat().Count()
	if err != nil {
		clog.Error("%s count error: %v, req: %v", fun, err, statListReq)
		stat.ReplyFail(w)
		return
	}

	stats, err := model.NewStat().List(statListReq.Offset, statListReq.Limit)
	if err != nil {
		clog.Error("%s list error: %v, req: %v", fun, err, statListReq)
		stat.ReplyFail(w)
		return
	}

	stat.ReplyOk(w, map[string]interface{}{
		"list":   stats,
		"total":  n,
		"offset": statListReq.Offset + len(stats),
	})
	return
}
