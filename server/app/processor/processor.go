package processor

import (
	"github.com/wonderivan/logger"
	"sync"
)

const (
	ChanBufferMaxLength = 128
	WarningChanBufferLeagth = ChanBufferMaxLength/2
)

type Processor struct {
	lock sync.Mutex
	Name string
	Closed 		bool
	MessageChan chan []byte
	StopChan	chan struct{}
}

func NewProcessor(name string) *Processor {
	return &Processor{
		Name:name,
		MessageChan:make(chan []byte , ChanBufferMaxLength),
		StopChan:make(chan struct{}),
	}
}

func (p *Processor) Receive(pattern string , message []byte){
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.Closed {
		return
	}
	defer func() {
		if r := recover() ; r!= nil {
			logger.Error("processor receive message failed : %s",r)
		}
	}()

	msg ,err := Encode(pattern,message)
	if err != nil {
		logger.Error("encode failed : %s",err)
	}

	count := len(p.MessageChan)
	if count < ChanBufferMaxLength {
		p.MessageChan <- msg
		count += 1
	}

	if count > WarningChanBufferLeagth {
		logger.Warn("processor message chan count warning! : %d",count)
	}
}

func (p *Processor) HandleLoop(handler func(pattern string ,message []byte)){
	for {
		select {
		case data := <- p.MessageChan:
			if data == nil {
				continue
			}
			pattern , msg ,err := Decode(data)
			if err != nil {
				logger.Error("processor loop handle ")
				continue
			}
			handler(pattern,msg)
		case <- p.StopChan:
			p.lock.Lock()
			defer p.lock.Unlock()
			close(p.MessageChan)
			p.Closed = true
			return
		}
	}
}

func (p *Processor) Stop(){
	close(p.StopChan)
}
