package visitor

import "sync"

type Visitors struct {
	lock 		sync.RWMutex

	Visitors 	map[string] *Visitor
}

var R = &Visitors{
	Visitors:make(map[string] *Visitor),
}

func Find(playerId string) *Visitor {
	R.lock.RLock()
	defer R.lock.RUnlock()

	return R.Visitors[playerId]
}

func Registry(v *Visitor) {
	R.lock.Lock()
	defer R.lock.Unlock()

	R.Visitors[v.PlayerId] = v
}

func UnRegistry(id string) {
	R.lock.Lock()
	defer R.lock.Unlock()

	delete(R.Visitors,id)
}