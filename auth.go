package core

import "errors"

var UserStatuses = newUserStatusRegistry()

type UserStatus struct {
	Name string
}

type userStatusRegistry struct {
	Registration *UserStatus
	Blocked      *UserStatus
	Disabled     *UserStatus
	Enabled      *UserStatus

	statuses []*UserStatus
}

func newUserStatusRegistry() *userStatusRegistry {

	registration := &UserStatus{Name: "registration"}
	blocked := &UserStatus{Name: "blocked"}
	disabled := &UserStatus{Name: "disabled"}
	enabled := &UserStatus{Name: "enabled"}

	return &userStatusRegistry{
		Registration: registration,
		Blocked:      blocked,
		Disabled:     disabled,
		Enabled:      enabled,

		statuses: []*UserStatus{registration, blocked, disabled, enabled},
	}
}

func (s *userStatusRegistry) List() []*UserStatus {
	return s.statuses
}

func (s *userStatusRegistry) Parse(status string) (*UserStatus, error) {

	for _, st := range s.List() {
		if st.Name == status {
			return st, nil
		}
	}

	return nil, errors.New("cloud not parse user status")
}
