package event

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/stan.go"
)

type Subscription interface {
	parseMsg(msg *stan.Msg, pb *proto.Message) error
	Listen() (stan.Subscription, error)
}

type subscription struct {
	Subject        Subject
	QueueGroupName QueueGroup
	OnMessage      func()
	Client         stan.Conn
	ProtoMsg       *proto.Message
	//AckWait int64
	//subOption stan.SubscriptionOption
}

func (s *subscription) Listen() (stan.Subscription, error) {
	return s.Client.QueueSubscribe(s.Subject.name, s.QueueGroupName.name, func(msg *stan.Msg) {
		if er := s.parseMsg(msg, s.ProtoMsg); er != nil {

		}
	}, stan.DurableName(s.QueueGroupName.name), stan.DeliverAllAvailable())
}

func (s *subscription) parseMsg(msg *stan.Msg, pb *proto.Message) error {
	return proto.Unmarshal(msg.Data, *pb)
}

func NewSubscription(subject Subject, group QueueGroup, client stan.Conn, msg *proto.Message) Subscription {
	return &subscription{Subject: subject, QueueGroupName: group, Client: client, ProtoMsg: msg}
}
