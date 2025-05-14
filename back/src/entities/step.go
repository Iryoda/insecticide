package entities

import (
	"time"

	"github.com/iryoda/insecticide/src/dtos"
)

type ActionType string

type AssertionType string

const (
	// Action
	CLICK       ActionType = "click"
	DOUBLECLICK ActionType = "doubleclick"

	// Input
	INPUT_TEXT ActionType = "text"

	// Window
	GOTO ActionType = "goto"
)

const (
	TEXT AssertionType = "text"
)

type Step struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	IsUnique        bool   `json:"isUnique" db:"isUnique"`
	Selector        string
	Action          *ActionType    `json:"action"`
	Assertion       *AssertionType `json:"assertion"`
	AssertionTarget *string        `json:"assertionTarget" db:"assertionTarget"`
	UpdatedAt       *time.Time     `json:"updatedAt" db:"updatedAt"`
	CreatedAt       time.Time      `json:"createdAt" db:"createdAt"`
}

func CreateStep(s dtos.CreateStep) (Step, error) {
	if err := s.Validate(); err != nil {
		return Step{}, err
	}

	step := Step{
		Name:   s.Name,
		Action: (*ActionType)(&s.Action),
	}

	if s.IsUnique != nil {
		step.IsUnique = *s.IsUnique
	}

	if s.Selector != nil {
		step.Selector = *s.Selector
	}

	if s.ComponentId != nil {
		step.Id = *s.ComponentId
	}

	if s.NextAssertId != nil {
		step.Id = *s.NextAssertId
	}

	return step, nil
}
