package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Adwait-aayush/ECMCP/internal/config"
	"github.com/Adwait-aayush/ECMCP/internal/database"
	"github.com/Adwait-aayush/ECMCP/internal/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database instance")
	}
	defer mainDB.Close()
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msgf("Starting server on port %s in %s mode", cfg.Server.Port, cfg.Server.GinMode)
}
