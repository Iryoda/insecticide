package repositories

import (
	"github.com/iryoda/insecticide/src/entities"
	"github.com/jmoiron/sqlx"
)

type StepRepository struct {
	db *sqlx.DB
}

func NewStepRepository(db *sqlx.DB) StepRepository {
	return StepRepository{
		db: db,
	}
}

func (sr StepRepository) Save(step entities.Step) (entities.Step, error) {
	stmt, err := sr.db.PrepareNamed(`INSERT INTO steps (
            name,
            isUnique,
            selector,
            action,
            assertion,
            assertionTarget,
            createdAt,
            updatedAt
        ) VALUES (
            :name, 
            :isUnique,
            :selector,
            :action,
            :assertion,
            :assertionTarget,
            :createdAt,
            :updatedAt
        ) RETURNING id`)
	if err != nil {
		return entities.Step{}, err
	}

	var id string
	err = stmt.Get(&id, step)
	if err != nil {
		return entities.Step{}, err
	}

	step, err = sr.GetById(id)
	if err != nil {
		return entities.Step{}, err
	}

	return step, nil
}

func (sr StepRepository) GetById(id string) (entities.Step, error) {
	var step entities.Step
	if err := sr.db.Get(&step, `SELECT * FROM steps WHERE id = $1`, id); err != nil {
		return entities.Step{}, err
	}

	return step, nil
}
