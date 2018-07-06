package service

import (
	zmq "github.com/pebbe/zmq4"
)

type _Publisher struct {
	msgChan chan string
	socket  *zmq.Socket
}

var publishers = make(map[string]*_Publisher)

func RunPublisher(_name string, _proc string) {
	socket, err := zmq.NewSocket(zmq.PUB)
	if nil != err {
		panic(err)
	}
	defer socket.Close()

	publisher := &_Publisher{}
	publisher.socket = socket
	publisher.socket.Bind(_proc)
	publisher.msgChan = make(chan string)
	publishers[_name] = publisher

	for {
		select {
		case msg := <-publisher.msgChan:
			publisher.socket.Send(msg, 0)
		}
	}

}

func Publish(_name string, _msg string) {
	if publisher, ok := publishers[_name]; !ok {
		publisher.msgChan <- _msg
	}
}
