package services

import (
	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/entities"
	"github.com/iryoda/insecticide/src/repositories"
	"github.com/iryoda/insecticide/src/utils"
)

type SuiteService struct {
	TestRepository repositories.SuiteRepository
}

func NewTestService(tr repositories.SuiteRepository) SuiteService {
	return SuiteService{
		TestRepository: tr,
	}
}

func (ts SuiteService) Create(data dtos.CreateSuite) (entities.Suite, error) {
	t, err := entities.NewTest(data)
	if err != nil {
		return entities.Suite{}, utils.ValidationError(err.Error())
	}

	if _, err = ts.TestRepository.GetByName(data.Name); err != nil {
		return entities.Suite{}, utils.ValidationError("There is already a test with the given name.")
	}

	test, err := ts.TestRepository.Save(t)
	if err != nil {
		return entities.Suite{}, err
	}

	return test, nil
}

func (ts SuiteService) GetById(id string) (entities.Suite, error) {
	test, err := ts.TestRepository.GetById(id)
	if err != nil {
		return entities.Suite{}, utils.NotFoundError("Test not found.")
	}

	return test, nil
}
