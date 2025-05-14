package services

import (
	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/entities"
	"github.com/iryoda/insecticide/src/repositories"
	"github.com/iryoda/insecticide/src/utils"
)

type TestComponentNodeService struct {
	TestComponentNodeRepository repositories.SuiteComponentNodeRepository
}

func NewTestComponentNodeService(tr repositories.SuiteComponentNodeRepository) TestComponentNodeService {
	return TestComponentNodeService{
		TestComponentNodeRepository: tr,
	}
}

func (ts TestComponentNodeService) Create(data dtos.CreateTestComponentNode) (entities.SuiteComponentNode, error) {
	t, err := entities.NewTestComponentNode(data)
	if err != nil {
		return entities.SuiteComponentNode{}, utils.ValidationError(err.Error())
	}

	test, err := ts.TestComponentNodeRepository.Save(t)
	if err != nil {
		return entities.SuiteComponentNode{}, err
	}

	return test, nil
}

func (ts TestComponentNodeService) GetById(id string) (entities.SuiteComponentNode, error) {
	test, err := ts.TestComponentNodeRepository.GetById(id)
	if err != nil {
		return entities.SuiteComponentNode{}, utils.NotFoundError("Test not found")
	}

	return test, nil
}
