package service

import (
	zmq "github.com/pebbe/zmq4"
)

type Broker struct {
	frontend *zmq.Socket
	backend  *zmq.Socket
}

func NewBroker() (*Broker, error) {
	frontend, err := zmq.NewSocket(zmq.ROUTER)
	if nil != err {
		return nil, err
	}

	backend, err := zmq.NewSocket(zmq.DEALER)
	if nil != err {
		frontend.Close()
		return nil, err
	}

	broker := &Broker{
		frontend: frontend,
		backend:  backend,
	}
	return broker, nil
}

func (this *Broker) Run(_procFrontend string, _procBackend string) {
	defer this.frontend.Close()
	defer this.backend.Close()

	this.frontend.Bind(_procFrontend)
	this.backend.Bind(_procBackend)

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(this.frontend, zmq.POLLIN)
	poller.Add(this.backend, zmq.POLLIN)
	//  Switch messages between sockets
	for {
		sockets, err := poller.Poll(-1)
		if nil != err {
			panic(err)
		}
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case this.frontend:
				for {
					data, _ := s.RecvBytes(0)
					more, _ := s.GetRcvmore()

					if more {
						this.backend.SendBytes(data, zmq.SNDMORE)
					} else {
						this.backend.SendBytes(data, 0)
						break
					}
				}
			case this.backend:
				for {
					data, _ := s.RecvBytes(0)
					more, _ := s.GetRcvmore()

					if more {
						this.frontend.SendBytes(data, zmq.SNDMORE)
					} else {
						this.frontend.SendBytes(data, 0)
						break
					}
				}
			}
		}
	}
}
