package repositories

import (
	"github.com/iryoda/insecticide/src/entities"
	"github.com/jmoiron/sqlx"
)

type SuiteRepository struct {
	db *sqlx.DB
}

func NewTestRepository(db *sqlx.DB) SuiteRepository {
	return SuiteRepository{
		db: db,
	}
}

func (tr SuiteRepository) Save(data entities.Suite) (entities.Suite, error) {
	if _, err := tr.db.NamedExec(
		`INSERT INTO test (
			name,
			selector,
			createdAt,
			updatedAt
		) VALUES (
			:name, 
			:selector,
			:createdAt,
			:updatedAt
		)`, data); err != nil {
		return entities.Suite{}, err
	}

	data, err := tr.GetById(data.Id)
	if err != nil {
		return entities.Suite{}, err
	}

	return data, nil
}

func (tr SuiteRepository) GetById(id string) (entities.Suite, error) {
	var test entities.Suite
	if err := tr.db.Get(&test, `SELECT * FROM test WHERE id = $1`, id); err != nil {
		return entities.Suite{}, err
	}

	return test, nil
}

func (tr SuiteRepository) GetByName(name string) (entities.Suite, error) {
	var test entities.Suite
	if err := tr.db.Get(&test, `SELECT * FROM test WHERE name = $1`, name); err != nil {
		return entities.Suite{}, err
	}

	return test, nil
}
