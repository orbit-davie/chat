package cache

import (
	"chat_server/app/processor"
	"github.com/wonderivan/logger"
	"sync"
)

type VisitorsManage struct {
	lock sync.RWMutex

	Visitors	map[string] processor.Actor
}

var Visitors = &VisitorsManage{
	Visitors:make(map[string] processor.Actor),
}

func Registry(playerId string , visitor processor.Actor) {
	Visitors.lock.Lock()
	defer Visitors.lock.Unlock()
	if _ , ok := Visitors.Visitors[playerId] ; ok {
		logger.Error("playerId is exist...")
		return
	}
	Visitors.Visitors[playerId] = visitor
}

func UnRegistry(playerId string) {
	Visitors.lock.Lock()
	defer Visitors.lock.Unlock()

	delete(Visitors.Visitors , playerId)
}

func Find(playerId string) processor.Actor {
	Visitors.lock.RLock()
	defer Visitors.lock.RUnlock()

	return Visitors.Visitors[playerId]
}

func AllVisitors() []processor.Actor {
	Visitors.lock.Lock()
	visitors := make([]processor.Actor , 0)

	for _ , v := range Visitors.Visitors {
		visitors = append(visitors , v)
	}
	Visitors.lock.Unlock()
	return visitors
}