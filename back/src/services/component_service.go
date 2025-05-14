package services

import (
	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/entities"
	"github.com/iryoda/insecticide/src/repositories"
	"github.com/iryoda/insecticide/src/utils"
)

type ComponentService struct {
	ComponentRepository repositories.ComponentRepository
}

func NewComponentService(cr repositories.ComponentRepository) ComponentService {
	return ComponentService{
		ComponentRepository: cr,
	}
}

func (cs ComponentService) Create(req dtos.CreateComponent) (entities.Component, error) {
	c, err := entities.NewComponent(req)
	if err != nil {
		return entities.Component{}, utils.ValidationError(err.Error())
	}

	c, err = cs.ComponentRepository.Save(c)
	if err != nil {
		return entities.Component{}, err
	}

	return c, nil
}

func (cs ComponentService) GetById(id string) (entities.Component, error) {
	component, err := cs.ComponentRepository.GetById(id)

	if err != nil {
		return entities.Component{}, utils.NotFoundError("Component not found")
	}

	return component, nil
}
