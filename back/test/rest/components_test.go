package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iryoda/insecticide/src/dtos"
)

type ComponentsSuite struct {
}

func ComponentsSetup() *ComponentsSuite {
	return &ComponentsSuite{}
}

func (s *ComponentsSuite) TearDown() {
}

func (s *ComponentsSuite) TestCreateComponent(t *testing.T) {
	description := "test"
	request := dtos.CreateComponent{
		Name:        "test",
		Description: &description,
	}
	json, _ := json.Marshal(request)

	httptest.NewRequest(http.MethodPost, "/components", bytes.NewBuffer(json))
}

func TestComponents(t *testing.T) {
	s := ComponentsSetup()

	t.Run("when create component, returns component", s.TestCreateComponent)
}
