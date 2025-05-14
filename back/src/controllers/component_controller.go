package controllers

import (
	"net/http"

	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/services"
	"github.com/iryoda/insecticide/src/utils"
	"github.com/labstack/echo/v4"
)

type ComponentController struct {
	ComponentService services.ComponentService
}

func NewComponentController(cs services.ComponentService) ComponentController {
	return ComponentController{
		ComponentService: cs,
	}
}

func (sc *ComponentController) Create(e echo.Context) error {
	var request dtos.CreateComponent
	err := e.Bind(&request)
	if err != nil {
		return utils.InvalidRequest(err.Error())
	}

	s, err := sc.ComponentService.Create(request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, s)
}

func (sc *ComponentController) GetById(e echo.Context) error {
	id := e.Param("id")

	s, err := sc.ComponentService.GetById(id)

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, s)
}
