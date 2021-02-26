package event

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/stan.go"
	"log"
)

type Publisher interface {
	parseMsg(msg *proto.Message) ([]byte, error)
	Publish() error
}

type publisher struct {
	Subject  Subject
	Client   stan.Conn
	ProtoMsg *proto.Message
	//AckWait int64
	//subOption stan.SubscriptionOption
}

func (p publisher) parseMsg(msg *proto.Message) ([]byte, error) {

	return proto.Marshal(*msg)
}

func (p publisher) Publish() error {
	d, er := p.parseMsg(p.ProtoMsg)
	if er != nil {
		log.Printf("Could not Parse protobuf message: %v", er)
		return er
	}

	return p.Client.Publish(p.Subject.name, d)
}

func NewPublisher(subject Subject, client stan.Conn, msg *proto.Message) Publisher {
	return &publisher{Subject: subject, Client: client, ProtoMsg: msg}
}
