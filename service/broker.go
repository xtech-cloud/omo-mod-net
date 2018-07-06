package service

import (
	zmq "github.com/pebbe/zmq4"
)

func RunBroker(_name string, _procFrontend string, _procBackend string) {
	frontend, err := zmq.NewSocket(zmq.ROUTER)
	if nil != err {
		panic(err)
	}
	defer frontend.Close()

	backend, err := zmq.NewSocket(zmq.DEALER)
	if nil != err {
		panic(err)
	}
	defer backend.Close()

	frontend.Bind(_procFrontend)
	backend.Bind(_procBackend)

	//  Initialize poll set
	poller := zmq.NewPoller()
	poller.Add(frontend, zmq.POLLIN)
	poller.Add(backend, zmq.POLLIN)

	//  Switch messages between sockets
	for {
		sockets, err := poller.Poll(-1)
		if nil != err {
			panic(err)
		}
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case frontend:
				for {
					data, _ := s.RecvBytes(0)
					more, _ := s.GetRcvmore()

					if more {
						backend.SendBytes(data, zmq.SNDMORE)
					} else {
						backend.SendBytes(data, 0)
						break
					}
				}
			case backend:
				for {
					data, _ := s.RecvBytes(0)
					more, _ := s.GetRcvmore()

					if more {
						frontend.SendBytes(data, zmq.SNDMORE)
					} else {
						frontend.SendBytes(data, 0)
						break
					}
				}
			}
		}
	}
}
