package entities

import (
	"time"

	"github.com/iryoda/insecticide/src/dtos"
)

type Component struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	HeadStepId  *string    `json:"headStepId"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func NewComponent(data dtos.CreateComponent) (Component, error) {
	if err := data.Validate(); err != nil {
		return Component{}, err
	}

	return Component{
		Name:        data.Name,
		Description: data.Description,
	}, nil
}
