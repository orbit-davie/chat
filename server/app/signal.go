package app

type Signal struct {
	Stop chan struct{}
	Stopped chan struct{}
}

func NewSignal() *Signal {
	return &Signal{
		Stop:make(chan struct{}),
		Stopped:make(chan struct{}),
	}
}
