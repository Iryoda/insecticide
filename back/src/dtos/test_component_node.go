package dtos

import "errors"

type CreateTestComponentNode struct {
	TestId          string  `json:"testId"`
	ComponentId     string  `json:"componentId"`
	NextComponentId *string `json:"nextComponentId"`
}

func (c CreateTestComponentNode) Validate() error {
	if c.TestId == "" {
		return errors.New("A TestId must be provided.")
	}

	if c.ComponentId == "" {
		return errors.New("A ComponentId must be provided.")
	}

	return nil
}
