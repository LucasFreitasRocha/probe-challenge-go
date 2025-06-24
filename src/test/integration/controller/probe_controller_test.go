package controller_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/routes"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateProbe(t *testing.T) {
	fmt.Println("Iniciando teste de criação de sonda")
	db, cleanup, err := SetupTestDB(context.Background())
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados: %v", err)
	}
	defer cleanup()
	db.AutoMigrate(&model.Probe{})
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



func SetupTestDB(ctx context.Context) (*gorm.DB, func(), error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "probe_test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, nil, err
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=test password=test dbname=probe_test sslmode=disable",
		host, port.Port())

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, err
	}

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			log.Printf("Erro ao parar o container: %v", err)
		}
	}
	

	return db, cleanup, nil
}

