package entities

import (
	"time"

	"github.com/iryoda/insecticide/src/dtos"
)

type Suite struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	HeadComponentId string     `json:"headComponentId"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	CreatedAt       time.Time  `json:"createdAt"`
}

func NewTest(data dtos.CreateSuite) (Suite, error) {
	if err := data.Validate(); err != nil {
		return Suite{}, err
	}

	return Suite{
		Name:            data.Name,
		HeadComponentId: data.HeadComponentId,
	}, nil
}
