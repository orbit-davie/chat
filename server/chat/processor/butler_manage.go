package processor

import (
	"chat_server/app/processor"
	"github.com/wonderivan/logger"
	"sync"
)

type ButlersManage struct {
	lock sync.RWMutex

	Butlers 	map[string] processor.Actor
}

var Butlers = &ButlersManage{
	Butlers:make(map[string] processor.Actor),
}

func Registry(playerId string  , butler processor.Actor) {
	Butlers.lock.Lock()
	defer Butlers.lock.Unlock()
	if _ , ok := Butlers.Butlers[playerId] ; ok {
		logger.Error("playerId is exist...")
		return
	}
	Butlers.Butlers[playerId] = butler
}

func UnRegistry(playerId string) {
	Butlers.lock.Lock()
	defer Butlers.lock.Unlock()

	delete(Butlers.Butlers,playerId)
}

func FindButler(playerId string) processor.Actor {
	Butlers.lock.RLock()
	defer Butlers.lock.RUnlock()

	return Butlers.Butlers[playerId]
}

func AllButlers() []processor.Actor {
	Butlers.lock.Lock()
	butlers := make([]processor.Actor , 0)

	for _ , v := range Butlers.Butlers {
		butlers = append(butlers , v)
	}
	Butlers.lock.Unlock()
	return butlers
}
