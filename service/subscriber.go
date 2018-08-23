package service

import (
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Subscriber struct {
	socket *zmq.Socket
}

func NewSubscriber() (*Subscriber, error) {
	socket, err := zmq.NewSocket(zmq.SUB)
	if nil != err {
		return nil, err
	}
	subscriber := &Subscriber{
		socket: socket,
	}
	return subscriber, nil
}

func (this *Subscriber) Run(_proc string, _filter string, _processor func(string)) {
	defer this.socket.Close()
	this.socket.Connect(_proc)
	this.socket.SetSubscribe(_filter)

	for {
		msg, _ := this.socket.Recv(0)

		if msgs := strings.Fields(msg); len(msgs) > 1 {
			if nil != _processor {
				_processor(msgs[1])
			}
		}
	}
}
