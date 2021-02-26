package event

var Subjects = newSubjectRegistry()

type Subject struct {
	name string
}

type subjectRegistry struct {
	AuthChange *Subject
}

func newSubjectRegistry() *subjectRegistry {
	return &subjectRegistry{
		AuthChange: &Subject{name: "auth:change"},
	}
}
