package lander

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Models interface {
	PK() bson.M
	Collection() string
	Index() []mgo.Index		//索引
}