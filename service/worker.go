package service

import (
	zmq "github.com/pebbe/zmq4"
)

type Worker struct {
	socket *zmq.Socket
}

func NewWorker() (*Worker, error) {
	socket, err := zmq.NewSocket(zmq.REP)
	if nil != err {
		panic(err)
	}
	worker := &Worker{
		socket: socket,
	}
	return worker, nil
}
func (this *Worker) Run(_proc string, _processor func([]byte) ([]byte, error)) {
	defer this.socket.Close()
	this.socket.Connect(_proc)

	for {
		//  Wait for next request from client
		req, err := this.socket.RecvBytes(0)
		if nil != err {
			continue
		}

		rsp, err := _processor(req)
		if nil != err {
			rsp = make([]byte, 1)
			rsp[0] = 255
			this.socket.SendBytes(rsp, 0)
			continue
		}
		this.socket.SendBytes(rsp, 0)
	}
}
