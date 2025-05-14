package api

import (
	"github.com/iryoda/insecticide/src/controllers"
	"github.com/iryoda/insecticide/src/repositories"
	"github.com/iryoda/insecticide/src/routers"
	"github.com/iryoda/insecticide/src/services"
	"github.com/iryoda/insecticide/src/utils"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func NewServer(db *sqlx.DB) *echo.Echo {
	// Repositories
	stepRepository := repositories.NewStepRepository(db)
	testRepository := repositories.NewTestRepository(db)
	componentRepository := repositories.NewComponentRepository(db)

	// Services
	stepService := services.NewStepService(stepRepository)
	testService := services.NewTestService(testRepository)
	componentService := services.NewComponentService(componentRepository)

	//  Controllers
	stepController := controllers.NewStepController(stepService)
	testController := controllers.NewTestController(testService)
	componentController := controllers.NewComponentController(componentService)

	// Routers
	stepRouter := routers.NewStepRouter(stepController)
	componentRouter := routers.NewComponentRouter(componentController)
	testRouter := routers.NewTestRouter(testController)

	e := echo.New()

	stepRouter.Register(e)
	componentRouter.Register(e)
	testRouter.Register(e)

	e.Use(utils.ErrorCatcherMiddleware)

	return e
}
