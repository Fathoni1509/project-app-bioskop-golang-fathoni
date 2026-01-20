package main

import (
	"log"
	"project-app-bioskop-golang-fathoni/cmd"
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/wire"
	"project-app-bioskop-golang-fathoni/pkg/database"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}

	logger, err := utils.InitLogger(config.PathLogging, config.Debug)

	repo := repository.NewRepository(db, logger)

	app := wire.Wiring(&repo, logger, config)

	cmd.ApiServer(app)
}