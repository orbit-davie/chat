package mongo

import (
	"chat_server/app/lander"
	"github.com/wonderivan/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoTask struct{
	collection 	string
	pattern 	string //模式
	doc 		bson.M
	change 		mgo.Change
	pk 			bson.M	//编译后的数据（数据对象 unmarshal）
}

type MongoManger struct {
	Conf		Config
	Mongo 		*Mongo
	Done 		chan *MongoTask
}

const MongoDbMaxTaskNum = 20480

var Default *MongoManger

func NewMongoManger(config Config , mgoConfig MongoCfg) *MongoManger {
	mgoSession , err := NewMongodb(mgoConfig)
	if err != nil {
		logger.Painc("mongodb session init failed!")
	}

	M := &MongoManger{
		Conf:config,
		Mongo:mgoSession,
		Done:make(chan *MongoTask , MongoDbMaxTaskNum),
	}
	if M.Mongo == nil {
		logger.Painc("mongodb is nil!")
	}
	return M
}

func(m *MongoManger) GoRun(t *MongoTask) {
	m.Done <- t
}

func (m *MongoManger) Land(models lander.Models) error {
	d , err := marshal(models)
	if err != nil {
		logger.Error("models marshal bson.M failed : %s", err)
		return err
	}
	task := &MongoTask{
		collection:models.Collection(),
		pattern:"upsert",
		doc:d,
		pk:models.PK(),
	}
	if taskNum := len(m.Done) ; taskNum >= MongoDbMaxTaskNum {
		go m.GoRun(task)
		return nil
	}
	m.Done <- task
	return nil
}

func (m *MongoManger) Remove(models lander.Models) error {
	task := &MongoTask{
		collection:models.Collection(),
		pattern:"remove",
		pk:models.PK(),
	}
	if taskNum := len(m.Done) ; taskNum >= MongoDbMaxTaskNum {
		go m.GoRun(task)
		return nil
	}
	m.Done <- task
	return nil
}

func (m *MongoManger) Run() {
	for {
		select {
		case task := <- m.Done:
			m.Mongo.MReturnChangeInfo(task.collection , func(collection *mgo.Collection) (changeInfo *mgo.ChangeInfo , err error) {
				switch task.pattern {
				case "upsert":
					changeInfo ,err = collection.Upsert(task.pk,task.doc)
				case "remove":
					err = collection.Remove(task.pk)
				}
				if err != nil {
					logger.Error("mongodb %s failed : %s" ,task.pattern, err)
				}
				return
			})
		}
		var output bool
		if remainCount := len(m.Done); remainCount < 10 {
			output = true
		} else if remainCount%1000 == 0 {
			output = true
		}
		if output {
			logger.Debug("landing: queue has `%d` remain", len(m.Done))
		}
	}
}

func(m *MongoManger) Stop(){
	for{
		if count := len(m.Done); count != 0 {
			logger.Info("landing: landing chan has %d items left", count)
			time.Sleep(time.Second)
		}else {
			return
		}
	}
}

func(m *MongoManger) StopSafe() {
	//处理完所有数据库写入操作
	m.Stop()
	m.Mongo.Close()
}

//将接口Model 编译为 bson.M ,
func marshal(models lander.Models) (doc bson.M , err error){
	data ,err := bson.Marshal(models)
	if err == nil {
		err = bson.Unmarshal(data , &doc)
	}
	return
}