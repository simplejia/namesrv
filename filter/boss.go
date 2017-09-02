package filter

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/simplejia/clog"
)

func Boss(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	err := m["__E__"]
	bt := m["__T__"].(time.Time)
	path := m["__P__"]

	if err != nil {
		clog.Error("Boss() path: %v, err: %v, stack: %s", path, err, debug.Stack())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		clog.Debug("Boss() path: %v, elapse: %s", path, time.Since(bt))
	}
	return true
}
