package core

var Subjects = newSubjectRegistry()

type Subject struct {
	Name string
}

type subjectRegistry struct {
	AuthChange            *Subject
	UserCreated           *Subject
	UserUpdated           *Subject
	UserArtistCreated     *Subject
	UserArtistUpdated     *Subject
	UserLabelCreated      *Subject
	UserLabelUpdated      *Subject
	UserPortfolioCreated  *Subject
	UserPortfolioUpdated  *Subject
	UserProductionCreated *Subject
	UserProductionUpdated *Subject
	UserCinemaCreated     *Subject
	UserCinemaUpdated     *Subject
}

func newSubjectRegistry() *subjectRegistry {
	return &subjectRegistry{
		AuthChange:            &Subject{Name: "auth:change"},
		UserCreated:           &Subject{Name: "user:user:created"},
		UserUpdated:           &Subject{Name: "user:user:updated"},
		UserArtistCreated:     &Subject{Name: "user:artist:created"},
		UserArtistUpdated:     &Subject{Name: "user:artist:updated"},
		UserLabelCreated:      &Subject{Name: "user:label:created"},
		UserLabelUpdated:      &Subject{Name: "user:label:updated"},
		UserPortfolioCreated:  &Subject{Name: "user:portfolio:created"},
		UserPortfolioUpdated:  &Subject{Name: "user:portfolio:updated"},
		UserProductionCreated: &Subject{Name: "user:production:created"},
		UserProductionUpdated: &Subject{Name: "user:production:updated"},
		UserCinemaCreated:     &Subject{Name: "user:cinema:created"},
		UserCinemaUpdated:     &Subject{Name: "user:cinema:updated"},
	}
}
