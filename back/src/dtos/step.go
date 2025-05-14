package dtos

import (
	"errors"
)

type CreateStep struct {
	Name            string  `json:"name"`
	Action          string  `json:"action"`
	IsUnique        *bool   `json:"isUnique" db:"isUnique"`
	Selector        *string `json:"selector"`
	ComponentId     *string `json:"componentId"`
	NextAssertId    *string `json:"nextAssertId"`
	Assertion       *string `json:"assertion"`
	AssertionTarget *string `json:"assertionTarget"`
}

type UpdateStep = CreateStep

func (s CreateStep) Validate() error {
	if s.Name == "" {
		return errors.New("A Name must be provided.")
	}

	if len(s.Name) > 255 {
		return errors.New("Name is too big.")
	}

	if s.Action == "" {
		return errors.New("An action must be provided.")
	}

	return nil
}
