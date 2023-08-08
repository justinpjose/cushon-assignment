package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinpjose/cushon-assignment/configs"
	"github.com/justinpjose/cushon-assignment/internal/db/postgres"
	"github.com/justinpjose/cushon-assignment/internal/logging"
	"github.com/justinpjose/cushon-assignment/internal/logging/zerolog"
	"github.com/justinpjose/cushon-assignment/internal/router/httprouter"
)

const (
	port       = 8080
	apiVersion = "v0"
	service    = "cushon-assignment"
)

var (
	inactiveTimeout = 20 * time.Second
)

func main() {
	log := zerolog.New()
	log.Field(logging.ServiceKey, service)
	log.Field(logging.APIVersionKey, apiVersion)

	dbCfg, err := configs.GetDBConfig()
	if err != nil {
		log.Fatalf("failed to get db configs: %v", err)
	}

	db, err := postgres.New(*dbCfg, log)
	if err != nil {
		log.Fatalf("failed to get postgres db: %v", err)
	}

	h := getHandlers(db, log)

	router := httprouter.New()
	getRoutes(router, h, apiVersion)

	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: inactiveTimeout,
		Handler:           router,
	}

	log.Infof("Listening on port %d", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Errorf("failed to serve requests : %v", err)
	}
}
