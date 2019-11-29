package processor

import (
	"chat_server/app/processor"
	"github.com/wonderivan/logger"
	"sync"
)

type RoomsManage struct {
	lock sync.RWMutex

	Rooms 		map[string] processor.Actor
}

var Rooms = &RoomsManage{
	Rooms:make(map[string] processor.Actor),
}

func RoomsRegistry(name string , room processor.Actor) {
	Rooms.lock.Lock()
	defer Rooms.lock.Unlock()
	if _ , ok := Rooms.Rooms[name] ; ok {
		logger.Error("playerId is exist...")
		return
	}
	Rooms.Rooms[name] = room
}

func RoomsUnRegistry(name string) {
	Rooms.lock.Lock()
	defer Rooms.lock.Unlock()

	delete(Rooms.Rooms , name)
}

func FindRoom(name string) processor.Actor {
	Rooms.lock.RLock()
	defer Rooms.lock.RUnlock()

	return Rooms.Rooms[name]
}

func RoomsAll() []processor.Actor {
	Rooms.lock.Lock()
	RoomsSlice := make([]processor.Actor , 0)

	for _ , v := range Rooms.Rooms {
		RoomsSlice = append(RoomsSlice , v)
	}
	Rooms.lock.Unlock()
	return RoomsSlice
}
