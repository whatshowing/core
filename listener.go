package core

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"github.com/nats-io/stan.go"
	"log"
)

type Listener interface {
	parseMsg(msg *stan.Msg, pb proto.Message) error
	Listen(onMsg func(pb proto.Message)) (stan.Subscription, error)
}

type listener struct {
	Subject        *Subject
	QueueGroupName *QueueGroup
	OnMessage      func()
	Client         stan.Conn
	ProtoMsg       proto.Message
	//AckWait int64
	//subOption stan.ListenerOption
}

func (s *listener) Listen(onMsg func(pb proto.Message)) (stan.Subscription, error) {
	return s.Client.QueueSubscribe(s.Subject.Name, s.QueueGroupName.Name, func(msg *stan.Msg) {
		en := protobuf.ProtobufEncoder{}

		if er := en.Decode(s.Subject.Name, msg.Data, s.ProtoMsg); er != nil {
			log.Printf("Error listening to subject %v. \n Error: %v \n", s.Subject.Name, er)
		}
		log.Printf("Listening event on subject %v\n", s.Subject.Name)
		onMsg(s.ProtoMsg)
		log.Println("event subject:", s.Subject.Name, " has been processed successfully")
	}, stan.DurableName(s.QueueGroupName.Name), stan.DeliverAllAvailable())
}

func (s *listener) parseMsg(msg *stan.Msg, pb proto.Message) error {
	return proto.Unmarshal(msg.Data, pb)
}

func NewListener(
	subject *Subject,
	group *QueueGroup,
	client stan.Conn,
	msg proto.Message,
) Listener {
	return &listener{
		Subject:        subject,
		QueueGroupName: group,
		Client:         client,
		ProtoMsg:       msg,
	}
}
