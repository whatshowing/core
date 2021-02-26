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
	d, er := e.Encode(p.Subject.Name, p.ProtoMsg)

	if er != nil {
		log.Printf("Could not Parse protobuf message: %v", er)
		return er
	}

	log.Printf("publishing to subject %v", p.Subject.Name)

	return p.Client.Publish(p.Subject.Name, d)
}

func NewPublisher(subject *Subject, client stan.Conn, msg proto.Message) Publisher {
	return &publisher{Subject: subject, Client: client, ProtoMsg: msg}
}
