package main

import (
	"errors"
	"flag"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ntnghiatn/financial-app-backend/internal/api"
	"github.com/ntnghiatn/financial-app-backend/internal/config"
	"github.com/ntnghiatn/financial-app-backend/internal/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	//
	log.SetLevel(log.DebugLevel)
	log.WithField("version", config.Version).Debug("starting Server....")

	// new instance database
	db, err := database.New()
	if err != nil {
		log.WithError(err).Fatal("could not get new instance database")
	}

	//
	router, err := api.NewRouter(db)
	if err != nil {
		log.WithError(err).Fatal("Error building router.")
	}

	//
	const addr = "0.0.0.0:8088"
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	//
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Error("Server failed!!")
	}

}
