package types

import "errors"

// Validate returns an error if the environment does not pass validation tests.
func (e *Environment) Validate() error {
	if err := ValidateName(e.Name); err != nil {
		return errors.New("environment name " + err.Error())
	}

	if len(e.Organization) == 0 {
		return errors.New("environment organization must be set")
	}

	return nil
}

// FixtureEnvironment returns a mocked environment.
func FixtureEnvironment(name string) *Environment {
	return &Environment{
		Name:         name,
		Organization: "default",
	}
}

// GetOrg gets the Organization that e belongs to.
func (e *Environment) GetOrg() string {
	return e.Organization
}

// GetEnv gets the Evironment that e belongs to (itself).
func (e *Environment) GetEnv() string {
	return e.Name
}
