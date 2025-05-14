package dtos

import "errors"

type CreateComponent struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (c CreateComponent) Validate() error {
	if c.Name == "" {
		return errors.New("A Name must be provided.")
	}

	if len(c.Name) > 255 {
		return errors.New("Name is too big.")
	}

	return nil
}
