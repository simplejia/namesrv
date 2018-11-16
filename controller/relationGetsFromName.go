package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/simplejia/clog/api"
	"github.com/simplejia/namesrv/model"
)

var AddNameStatFunc = func() func(string, string) error {
	fun := "AddNameStatFunc"
	ch := make(chan [2]string, 1e6)

	go func() {
		m := make(map[string]time.Time)
		n := make(map[string]int)
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				for key, expire := range m {
					if time.Since(expire) < 0 {
						continue
					}
					num := n[key]

					delete(m, key)
					delete(n, key)

					statModel := model.NewStat()
					statModel.Name = key
					err := statModel.IncNumDay(num)
					if err != nil {
						clog.Error("%s stat.IncNumDay err: %v, req: %v", fun, err, statModel)
					}
				}
			case kv := <-ch:
				key, value := kv[0], kv[1]
				num, _ := strconv.Atoi(value)
				num++
				if _, ok := n[key]; ok {
					n[key] += num
					break
				}
				m[key] = time.Now().Add(time.Second * 10)
				n[key] = num
			}
		}
	}()

	return func(key, value string) (err error) {
		select {
		case ch <- [2]string{key, value}:
		default:
		}
		return
	}
}()

// @postfilter("Boss")
func (relation *Relation) GetsFromName(w http.ResponseWriter, r *http.Request) {
	fun := "relation.GetsFromName"

	name := r.FormValue("name")
	cc := r.FormValue("cc")
	num := r.FormValue("num")

	AddNameStatFunc(name, num)

	relationModel := model.NewRelation()
	relationModel.Name = name

	rels, err := relationModel.GetsFromName()
	if err != nil {
		clog.Error("%s gets error: %v, req: %v", fun, err, relationModel)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ccNew := model.Relations(rels).CheckCode()
	if cc == ccNew {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"rels": rels,
		"cc":   ccNew,
	})
	return
}
