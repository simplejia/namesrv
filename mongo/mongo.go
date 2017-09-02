package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"github.com/simplejia/namesrv/conf"
	"github.com/simplejia/utils"
)

type Conf struct {
	ConnMaxLifetime string
	MaxIdleConns    int
	MaxOpenConns    int
	Dsn             string
}

var (
	DBS  map[string]*mgo.Session = map[string]*mgo.Session{}
	Envs map[string]*Conf
	Env  string
	C    *Conf
)

func parseDBFile(path string) {
	fcontent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fcontent = utils.RemoveAnnotation(fcontent)
	if err := json.Unmarshal(fcontent, &Envs); err != nil {
		panic(err)
	}

	Env = conf.Env
	C = Envs[Env]
	if C == nil {
		fmt.Println("env not right:", Env)
		os.Exit(-1)
	}

	session, err := mgo.Dial(C.Dsn)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.SecondaryPreferred, true)

	key := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	DBS[key] = session
}

func init() {
	dir := "mongo"
	for i := 0; i < 3; i++ {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			break
		}
		dir = filepath.Join("..", dir)
	}
	err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) (reterr error) {
			if err != nil {
				reterr = err
				return
			}
			if info.IsDir() {
				return
			}
			if filepath.Ext(path) != ".json" {
				return
			}

			parseDBFile(path)
			return
		},
	)
	if err != nil {
		panic(err)
	}
}
