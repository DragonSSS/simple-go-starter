package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DragonSSS/simple-go-starter/handlers"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// @title Simple Go Starter
// @version 1.0
// @description simple go starter template
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email scarecrow@gamil.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	router := mux.NewRouter()

	// healthz
	router.HandleFunc("/healthz", handlers.HealthCheck).Methods("GET")

	service := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Info("Server started")
	<-done
	log.Info("Server Stopping")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// some cleanup work, e.g db connection, data backup
		cancel()
	}()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Info("Server Exited Properly")
}
