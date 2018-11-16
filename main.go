// 名字服务(server).
// author: simplejia
// date: 2017/07/30

//go:generate wsp -l

package main

import (
	"fmt"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/lc"
	"github.com/simplejia/namecli/api"
	"github.com/simplejia/namesrv/conf"
	"github.com/simplejia/utils"

	"net/http"
)

func init() {
	lc.Init(1e5)

	clog.AddrFunc = func() (string, error) {
		return api.Name(conf.C.Addrs.Clog)
	}
	clog.Init(conf.C.App.Name, "", conf.C.Clog.Level, conf.C.Clog.Mode)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}

func main() {
	fun := "main"
	clog.Info("main()")

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.C.App.Port)
	err := utils.ListenAndServe(addr, nil)
	if err != nil {
		clog.Error("%s err: %v, addr: %v", fun, err, addr)
	}
}
