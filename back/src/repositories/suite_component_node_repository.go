package repositories

import (
	"github.com/iryoda/insecticide/src/entities"
	"github.com/jmoiron/sqlx"
)

type SuiteComponentNodeRepository struct {
	db *sqlx.DB
}

func NewTestComponentNodeRepository(db *sqlx.DB) SuiteComponentNodeRepository {
	return SuiteComponentNodeRepository{
		db: db,
	}
}

func (tr SuiteComponentNodeRepository) Save(test entities.SuiteComponentNode) (entities.SuiteComponentNode, error) {
	if _, err := tr.db.NamedExec(
		`INSERT INTO test_component_node (
			testId,
			componentId,
			nextComponentId,
		) VALUES (
			:testId 
			:componentId,
			:nextComponentId,
		)`, test); err != nil {
		return entities.SuiteComponentNode{}, err
	}

	test, err := tr.GetById(test.Id)
	if err != nil {
		return entities.SuiteComponentNode{}, err
	}

	return test, nil
}

func (tr SuiteComponentNodeRepository) GetById(id string) (entities.SuiteComponentNode, error) {
	var test entities.SuiteComponentNode
	if err := tr.db.Get(&test, `SELECT * FROM test_component_node WHERE id = $1`, id); err != nil {
		return entities.SuiteComponentNode{}, err
	}

	return test, nil
}
