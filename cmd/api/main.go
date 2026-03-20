package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Adwait-aayush/ECMCP/internal/config"
	"github.com/Adwait-aayush/ECMCP/internal/database"
	"github.com/Adwait-aayush/ECMCP/internal/logger"
	"github.com/Adwait-aayush/ECMCP/internal/server"
	"github.com/Adwait-aayush/ECMCP/internal/services"
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

	authService:=services.NewAuthService(db,cfg)
	productService:=services.NewProductService(db)
	userService:=services.NewUserService(db)
	srv := server.New(cfg, db, log,authService,productService,userService)
	router := srv.SetupRoute()

	httpServer := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Msgf("Starting server on port %s in %s mode", cfg.Server.Port, cfg.Server.GinMode)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown server")
	}
	log.Info().Msg("Server gracefully stopped")
}
