package entities

import "github.com/iryoda/insecticide/src/dtos"

type SuiteComponentNode struct {
	Id              string  `json:"id"`
	TestId          string  `json:"testId"`
	ComponentId     string  `json:"stepId"`
	NextComponentId *string `json:"nextComponentId"`
}

func NewTestComponentNode(data dtos.CreateTestComponentNode) (SuiteComponentNode, error) {
	if err := data.Validate(); err != nil {
		return SuiteComponentNode{}, err
	}

	return SuiteComponentNode{
		TestId:          data.TestId,
		ComponentId:     data.ComponentId,
		NextComponentId: data.NextComponentId,
	}, nil
}
