package event


var Subjects = newSubjectRegistry()

type subjectRegistry struct {
	AuthChange string
}


func newSubjectRegistry() *subjectRegistry {
	return &subjectRegistry{
		AuthChange: "auth:change",
	}
}

