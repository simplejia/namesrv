package controller

import (
	"net/http"

	cf "github.com/simplejia/namesrv/conf"
)

// @postfilter("Boss")
func (conf *Conf) Get(w http.ResponseWriter, r *http.Request) {
	cf.Cgi(w, r)
	return
}
