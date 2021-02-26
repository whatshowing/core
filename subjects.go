package core

var Subjects = newSubjectRegistry()

type Subject struct {
	Name string
}

type subjectRegistry struct {
	AuthChange  *Subject
	UserCreated *Subject
}

func newSubjectRegistry() *subjectRegistry {
	return &subjectRegistry{
		AuthChange:  &Subject{Name: "auth:change"},
		UserCreated: &Subject{Name: "user:created"},
	}
}
