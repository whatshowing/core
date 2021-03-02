package core

var UserStatuses = newSubjectRegistry()

type UserStatus struct {
	Name string
}

type userStatusRegistry struct {
	Registration *UserStatus
	Blocked      *UserStatus
	Disabled     *UserStatus
	Enabled      *UserStatus
}

func newUserStatusRegistry() *userStatusRegistry {
	return &userStatusRegistry{
		Registration: &UserStatus{Name: "registration"},
		Blocked:      &UserStatus{Name: "blocked"},
		Disabled:     &UserStatus{Name: "disabled"},
		Enabled:      &UserStatus{Name: "enabled"},
	}
}
