package core

var MicroServices = newMicroServiceRegistry()

type MicroService struct {
	Name string
}

type serviceRegistry struct {
	Auth      *MicroService
	Messaging *MicroService
	User      *MicroService
}

func newMicroServiceRegistry() *serviceRegistry {
	return &serviceRegistry{
		Auth:      &MicroService{Name: "auth"},
		Messaging: &MicroService{Name: "messaging"},
		User:      &MicroService{Name: "user"},
	}
}
