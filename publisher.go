package core

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"github.com/nats-io/stan.go"
	"log"
)

type Publisher interface {
	Publish() error
}

type publisher struct {
	Subject  *Subject
	Client   stan.Conn
	ProtoMsg proto.Message
	//AckWait int64
}

func (p publisher) Publish() error {
	e := protobuf.ProtobufEncoder{}
	log.Println("Proto Message: ", p.ProtoMsg)
	d, er := e.Encode(p.Subject.Name, p.ProtoMsg)

	log.Println("Bytes: ", d)
	if er != nil {
		log.Printf("Could not Parse protobuf message: %v", er)
		return er
	}

	log.Printf("publishing to subject %v", p.Subject.Name)

	err := p.Client.Publish(p.Subject.Name, d)

	if err != nil {
		log.Println(p.Subject.Name, " ", err)
	}

	log.Println("Published to subject:  ", p.Subject.Name)

	return nil
}

func NewPublisher(subject *Subject, client stan.Conn, msg proto.Message) Publisher {
	return &publisher{Subject: subject, Client: client, ProtoMsg: msg}
}
