package services

import (
	"fmt"

	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/entities"
	"github.com/iryoda/insecticide/src/repositories"
	"github.com/iryoda/insecticide/src/utils"
)

type StepService struct {
	StepRepository repositories.StepRepository
}

func NewStepService(sr repositories.StepRepository) StepService {
	return StepService{
		StepRepository: sr,
	}
}

func (s StepService) Create(req dtos.CreateStep) (entities.Step, error) {
	step, err := entities.CreateStep(req)
	if err != nil {
		return entities.Step{}, utils.ValidationError(err.Error())
	}

	created, err := s.StepRepository.Save(step)
	if err != nil {
		return entities.Step{}, err
	}

	fmt.Println("Step created:", created)
	return created, nil
}

func (s StepService) GetById(id string) (entities.Step, error) {
	step, err := s.StepRepository.GetById(id)

	if err == nil {
		return entities.Step{}, utils.NotFoundError("Step not found")
	}

	return step, nil
}
