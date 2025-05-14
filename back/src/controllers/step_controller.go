package controllers

import (
	"net/http"

	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/services"
	"github.com/iryoda/insecticide/src/utils"
	"github.com/labstack/echo/v4"
)

type StepController struct {
	StepService services.StepService
}

func NewStepController(stepService services.StepService) StepController {
	return StepController{
		StepService: stepService,
	}
}

func (sc *StepController) Create(e echo.Context) error {
	var request dtos.CreateStep
	err := e.Bind(&request)
	if err != nil {
		return utils.InvalidRequest(err.Error())
	}

	s, err := sc.StepService.Create(request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, s)
}

func (sc *StepController) GetById(e echo.Context) error {
	id := e.Param("id")

	s, err := sc.StepService.GetById(id)

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, s)
}
