package controllers

import (
	"net/http"

	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/services"
	"github.com/iryoda/insecticide/src/utils"
	"github.com/labstack/echo/v4"
)

type SuiteController struct {
	SuiteService services.SuiteService
}

func NewTestController(stepService services.SuiteService) SuiteController {
	return SuiteController{
		SuiteService: stepService,
	}
}

func (sc *SuiteController) Create(e echo.Context) error {
	var req dtos.CreateSuite
	err := e.Bind(&req)
	if err != nil {
		return utils.InvalidRequest(err.Error())
	}

	s, err := sc.SuiteService.Create(req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, s)
}

func (sc *SuiteController) GetById(e echo.Context) error {
	id := e.Param("id")

	s, err := sc.SuiteService.GetById(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, s)
}
