package controller

import (
	"net"
	"net/http"
	"strconv"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/namesrv/model"
)

// @postfilter("Boss")
func (relation *Relation) ReportOff(w http.ResponseWriter, r *http.Request) {
	fun := "relation.ReportOff"

	ipport := r.FormValue("ipport")
	ip, port, err := net.SplitHostPort(ipport)
	if err != nil {
		clog.Error("%s req error: %v, req: %v", fun, err, ipport)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_port, _ := strconv.Atoi(port)

	off := r.FormValue("off")
	_off, _ := strconv.ParseBool(off)

	relationModel := model.NewRelation()
	relationModel.Ip = ip
	relationModel.Port = _port
	relationModel.Off = _off

	err = relationModel.UpdateOff()
	if err != nil {
		clog.Error("%s update error: %v, req: %v", fun, err, relationModel)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
