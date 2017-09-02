// 名字服务(server).
// author: simplejia
// date: 2017/07/30

//go:generate wsp -l

package main

import (
	"fmt"

	"github.com/simplejia/clog"
	"github.com/simplejia/lc"
	"github.com/simplejia/namecli/api"
	"github.com/simplejia/namesrv/conf"
	_ "github.com/simplejia/namesrv/mongo"
	"github.com/simplejia/utils"

	"net/http"
)

func init() {
	lc.Init(1e5)

	clog.AddrFunc = func() (string, error) {
		return api.Name("clog.srv.ns")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}

func main() {
	clog.Info("main()")

	c := conf.Get()
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", c.App.Port)
	err := utils.ListenAndServe(addr, nil)
	if err != nil {
		clog.Error("main() err: %v", err)
	}
}
