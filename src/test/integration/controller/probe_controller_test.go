package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/rest_err"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/routes"
	"github.com/gin-gonic/gin"
)

func TestCreateProbeSuccess(t *testing.T) {
	db, cleanup, err := setupTestDB(context.Background())
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados: %v", err)
	}
	defer cleanup()
	
	r := gin.Default()
	probeService := config.InitProbeService(db)
	routes.SetupProbeRoutes(r, config.InitProbeController(probeService))
	body := []byte(`{"name":"Robo1","direction":"N","position_x":1,"position_y":2}`)
	req, _ := http.NewRequest("POST", "/probes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava 201, recebeu %d", w.Code)
	}
}

func TestCreateProbeInvalidDirection(t *testing.T) {
	db, cleanup, err := setupTestDB(context.Background())
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados: %v", err)
	}
	defer cleanup()
	r := gin.Default()
	probeService := config.InitProbeService(db)
	routes.SetupProbeRoutes(r, config.InitProbeController(probeService))
	body := []byte(`{"name":"Robo1","direction":"X","position_x":1,"position_y":2}`)
	req, _ := http.NewRequest("POST", "/probes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("esperava 400, recebeu %d", w.Code)
	}
	bodyBytes := w.Body.Bytes()

	var restErr rest_err.RestErr
	if err := json.Unmarshal(bodyBytes, &restErr); err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}
	var directionMsg string
	for _, cause := range restErr.Causes {
		if cause.Field == "direction" {
			directionMsg = cause.Message
			if directionMsg != "direction not valid, must be one of: N, E, S, W" {
				t.Fatalf("mensagem de erro inesperada: %s", directionMsg)
			}
		}
	}

}






