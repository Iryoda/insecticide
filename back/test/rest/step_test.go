package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iryoda/insecticide/src/api"
	"github.com/iryoda/insecticide/src/dtos"
	"github.com/iryoda/insecticide/src/entities"
	test_containers "github.com/iryoda/insecticide/test/containers"
	test_utils "github.com/iryoda/insecticide/test/utils"
	"github.com/labstack/echo/v4"
)

type StepSuite struct {
	T                 *testing.T
	PostgresContainer *test_containers.PostgresContainer
	Ctx               context.Context
	Server            *echo.Echo
}

func Setup(T *testing.T) *StepSuite {
	ctx := context.Background()
	pc := test_containers.NewPostgresTestContainer(ctx)
	e := api.NewServer(pc.DB)
	test_utils.SetRandomPortEcho(e)

	return &StepSuite{
		T:                 T,
		PostgresContainer: pc,
		Ctx:               ctx,
		Server:            e,
	}
}

func (s *StepSuite) TearDown() {
	s.PostgresContainer.Container.Terminate(s.Ctx)
}

func (s *StepSuite) TestCreateStep(t *testing.T) {
	request := dtos.CreateStep{
		Name:   "test",
		Action: string(entities.CLICK),
	}
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/steps", bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.Server.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())

	rec.Body.String()

	if rec.Code != http.StatusCreated {
		t.Fatalf("failed to create step: %v", rec.Body.String())
	}
}

func (s *StepSuite) TestCreateStepWithInvalidRequest(t *testing.T) {
	request := dtos.CreateStep{
		Name:   "test",
		Action: string(entities.CLICK),
	}
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/steps", bytes.NewBuffer(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.Server.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("failed to create step: %v", rec.Body.String())
	}
}

func TestStep(t *testing.T) {
	s := Setup(t)
	defer s.TearDown()

	t.Run("when create step, returns step", s.TestCreateStep)
	t.Run("when create step, returns bad request", s.TestCreateStepWithInvalidRequest)
}
