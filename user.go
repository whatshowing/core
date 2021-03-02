package core

import "errors"

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

	steps []*RegistrationStep
}

func newRegistrationStepRegistry() *registrationStepRegistry {

	initial := &RegistrationStep{Name: "initial"}
	phoneInput := &RegistrationStep{Name: "phone-input"}
	phoneVerification := &RegistrationStep{Name: "phone-verification"}
	personalNamesInput := &RegistrationStep{Name: "personal-names-input"}
	addressInput := &RegistrationStep{Name: "address-input"}
	socialLinksInput := &RegistrationStep{Name: "social-links-input"}
	usernameAvatarInput := &RegistrationStep{Name: "username-avatar-input"}
	physicalDescription := &RegistrationStep{Name: "physical-description"}
	characterSelection := &RegistrationStep{Name: "character-selection"}
	talentSelection := &RegistrationStep{Name: "talent-selection"}
	completed := &RegistrationStep{Name: "completed"}

	return &registrationStepRegistry{
		Initial:             initial,
		PhoneInput:          phoneInput,
		PhoneVerification:   phoneVerification,
		PersonalNamesInput:  personalNamesInput,
		AddressInput:        addressInput,
		SocialLinksInput:    socialLinksInput,
		UsernameAvatarInput: usernameAvatarInput,
		PhysicalDescription: physicalDescription,
		CharacterSelection:  characterSelection,
		TalentSelection:     talentSelection,
		Completed:           completed,

		steps: []*RegistrationStep{
			initial,
			phoneInput,
			phoneVerification,
			personalNamesInput,
			addressInput,
			socialLinksInput,
			usernameAvatarInput,
			physicalDescription,
			characterSelection,
			talentSelection,
			completed,
		},
	}

}

func (s *registrationStepRegistry) List() []*RegistrationStep {
	return s.steps
}

func (s *registrationStepRegistry) Parse(step string) (*RegistrationStep, error) {

	for _, sp := range s.List() {
		if sp.Name == step {
			return sp, nil
		}
	}

	return nil, errors.New("cloud not parse step")
}
