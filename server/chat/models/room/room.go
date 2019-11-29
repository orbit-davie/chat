package room

import (
	"github.com/satori/go.uuid"
)

const MaxVisitors = 4

type Room struct {
	MaxNum 		int32
	Name 		string
	Id 			string
	Visitors  	map[string] bool
}

func NewRoom(name string) *Room {
	return &Room{
		Name:name,
		Id:uuid.NewV4().String(),
		MaxNum:MaxVisitors,
		Visitors: make(map[string] bool),
	}
}

func (r *Room) Enter(vId string) bool {
	if _ , ok := r.Visitors[vId] ; ok {
		return false
	}
	if int32(len(r.Visitors)) == r.MaxNum {
		return false
	}
	r.Visitors[vId] = true
	return true
}

func (r *Room) Leave(vId string) bool {
	if _ , ok := r.Visitors[vId] ; !ok {
		return false
	}
	delete(r.Visitors , vId)
	return true
}

func (r *Room) All() (ids []string) {
	for id := range r.Visitors {
		ids = append(ids , id)
	}
	return
}

func (r *Room) Check(vId string) bool {
	var found bool
	for id := range r.Visitors {
		if id == vId {
			found = true
		}
	}
	return found
}