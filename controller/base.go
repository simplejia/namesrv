package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IBase interface {
	SetParam(string, interface{})
	GetParam(string) (interface{}, bool)
}

type Base struct {
	params map[string]interface{}
}

func (base *Base) SetParam(key string, value interface{}) {
	if base.params == nil {
		base.params = make(map[string]interface{})
	}
	base.params[key] = value
}

func (base *Base) GetParam(key string) (value interface{}, ok bool) {
	value, ok = base.params[key]
	return
}

func (base *Base) ReadBody(r *http.Request) (body []byte) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return
}

func (base *Base) ReplyOk(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ret":  1,
		"data": data,
	})
}

func (base *Base) ReplyFail(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ret": -1,
	})
}
