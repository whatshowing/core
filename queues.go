package core

var QueueGroups = newQueueGroupRegistry()

type QueueGroup struct {
	Name string
}

type queueGroupRegistry struct {
	AuthGroup      *QueueGroup
	MessagingGroup *QueueGroup
	UserGroup      *QueueGroup
}

func newQueueGroupRegistry() *queueGroupRegistry {
	return &queueGroupRegistry{
		AuthGroup:      &QueueGroup{Name: "auth"},
		MessagingGroup: &QueueGroup{Name: "messaging"},
		UserGroup:      &QueueGroup{Name: "user"},
	}
}
