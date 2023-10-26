package main

import (
	"context"
	"fmt"
	"jobnew/internal/auth"
	"jobnew/internal/database"
	"jobnew/internal/handlers"
	"jobnew/internal/models"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic() //.Err(err).Send()
	}
	// log.Info().Msg("hello this is our app")
}
func StartApp() error {
	privatePem, err := os.ReadFile("private.pem")
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}
	publicPem, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return fmt.Errorf("reading auth public key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPem)
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w ", err)
	}
	ms, err := models.NewService(db)
	if err != nil {
		return err
	}
	err = ms.AutoMigrate()
	if err != nil {
		return err
	}
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(a, ms),
	}
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			err = api.Close()
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil
}
