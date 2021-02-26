package core

var QueueGroups = newQueueGroupRegistry()

type QueueGroup struct {
	Name string
}

type queueGroupRegistry struct {
	AuthGroup *QueueGroup
	UserGroup *QueueGroup
}

func newQueueGroupRegistry() *queueGroupRegistry {
	return &queueGroupRegistry{
		AuthGroup: &QueueGroup{Name: "auth"},
		UserGroup: &QueueGroup{Name: "user"},
	}
}
