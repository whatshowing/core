package core

var QueueGroups = newQueueGroupRegistry()

type QueueGroup struct {
	name string
}

type queueGroupRegistry struct {
	AuthGroup *QueueGroup
}

func newQueueGroupRegistry() *queueGroupRegistry {
	return &queueGroupRegistry{
		AuthGroup: &QueueGroup{name: "auth"},
	}
}
