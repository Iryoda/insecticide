package repositories

import (
	"github.com/iryoda/insecticide/src/entities"
	"github.com/jmoiron/sqlx"
)

type ComponentRepository struct {
	db *sqlx.DB
}

func NewComponentRepository(db *sqlx.DB) ComponentRepository {
	return ComponentRepository{
		db: db,
	}
}

func (cr ComponentRepository) Save(component entities.Component) (entities.Component, error) {
	if _, err := cr.db.NamedExec(
		`INSERT INTO component (
			name,
			selector,
			createdAt,
			updatedAt
		) VALUES (
			:name, 
			:selector,
			:createdAt,
			:updatedAt
		)`, component); err != nil {
		return entities.Component{}, err
	}

	component, err := cr.GetById(component.Id)
	if err != nil {
		return entities.Component{}, err
	}

	return component, nil
}

func (cr ComponentRepository) GetById(id string) (entities.Component, error) {
	var component entities.Component
	if err := cr.db.Get(&component, `SELECT * FROM component WHERE id = $1`, id); err != nil {
		return entities.Component{}, err
	}

	return component, nil
}
