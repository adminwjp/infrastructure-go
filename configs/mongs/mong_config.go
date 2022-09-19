package mongs

import (
	"gopkg.in/mgo.v2"
)
var MongConfigInstance=&MongConfig{}
type MongConfig struct {
	Session *mgo.Session

}
