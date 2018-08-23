package service

import (
	zmq "github.com/pebbe/zmq4"
)

type Publisher struct {
	msgChan chan string
	socket  *zmq.Socket
}

func NewPublisher() (*Publisher, error) {
	socket, err := zmq.NewSocket(zmq.PUB)
	if nil != err {
		return nil, err
	}

	publisher := &Publisher{}
	publisher.socket = socket
	publisher.msgChan = make(chan string)
	return publisher, nil
}

func (this *Publisher) Run(_proc string) {
	defer this.socket.Close()
	this.socket.Bind(_proc)
	for {
		select {
		case msg := <-this.msgChan:
			this.socket.Send(msg, 0)
		}
	}
}

func (this *Publisher) Publish(_msg string) {
	this.msgChan <- _msg
}
