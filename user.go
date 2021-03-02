package core

var RegistrationSteps = newRegistrationStepRegistry()

type RegistrationStep struct {
	Name string
}

type registrationStepRegistry struct {
	Initial             *RegistrationStep
	PhoneInput          *RegistrationStep
	PhoneVerification   *RegistrationStep
	PersonalNamesInput  *RegistrationStep
	AddressInput        *RegistrationStep
	SocialLinksInput    *RegistrationStep
	UsernameAvatarInput *RegistrationStep
	PhysicalDescription *RegistrationStep
	CharacterSelection  *RegistrationStep
	TalentSelection     *RegistrationStep
	Completed           *RegistrationStep
}

func newRegistrationStepRegistry() *registrationStepRegistry {
	return &registrationStepRegistry{
		Initial:             &RegistrationStep{Name: "initial"},
		PhoneInput:          &RegistrationStep{Name: "phone-input"},
		PhoneVerification:   &RegistrationStep{Name: "phone-verification"},
		PersonalNamesInput:  &RegistrationStep{Name: "personal-names-input"},
		AddressInput:        &RegistrationStep{Name: "address-input"},
		SocialLinksInput:    &RegistrationStep{Name: "social-links-input"},
		UsernameAvatarInput: &RegistrationStep{Name: "username-avatar-input"},
		PhysicalDescription: &RegistrationStep{Name: "physical-description"},
		CharacterSelection:  &RegistrationStep{Name: "character-selection"},
		TalentSelection:     &RegistrationStep{Name: "talent-selection"},
		Completed:           &RegistrationStep{Name: "completed"},
	}
}
