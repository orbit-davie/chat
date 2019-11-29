package mongo

import (
	"github.com/wonderivan/logger"
	mongo"gopkg.in/mgo.v2"
)

const (
	URL = "127.0.0.1:27017" //mongodb连接字符串
	DBNAME = "chat"
)

type Mongo struct {
	Conf 		MongoCfg
	Session 	*mongo.Session
}

//Mongodb 基础方法封装

func NewMongodb(config MongoCfg) ( *Mongo , error){
	mgo := &Mongo{}
	mgo.Conf = config

	//session ,err := mongo.Dial(mgo.Conf.Url)
	info := &mongo.DialInfo{
		Username:"root",
		Password:"w55665978",
		Direct:false,
		Addrs: []string{mgo.Conf.Url},
		Source:"admin",
		Database:mgo.Conf.Database,
	}
	session , err := mongo.DialWithInfo(info)
	if err != nil {
		logger.Error("mongo Dial error : %s",err)
		return nil , err
	}
	mgo.Session = session
	logger.Info("dial mongodb server successfully!")
	return mgo , nil
}

func (m *Mongo ) M (collection string ,f func(*mongo.Collection) (err error)) error {
	session := m.Session.Clone()
	defer session.Close()
	c := session.DB(m.Conf.Database).C(collection)
	return f(c)
}

func (m *Mongo ) MReturnChangeInfo (collection string ,f func(*mongo.Collection) (*mongo.ChangeInfo , error)) (*mongo.ChangeInfo , error) {
	session := m.Session.Clone()
	defer session.Close()
	c := session.DB(m.Conf.Database).C(collection)

	return f(c)
}

//drop database
func (m *Mongo) DropDatabase() error {
	session := m.Session.Clone()
	defer session.Close()
	err := session.DB(m.Conf.Database).DropDatabase()
	return err
}

func (m *Mongo) Close() {
	if m.Session != nil {
		m.Session.Close()
	}
}
