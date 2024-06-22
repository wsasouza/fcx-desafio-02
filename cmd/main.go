package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/wsasouza/fcx-desafio-02/internal/infra/http"
	"github.com/wsasouza/fcx-desafio-02/internal/infra/repository"
	"github.com/wsasouza/fcx-desafio-02/internal/usecase"
)

func main() {
	repo, err := repository.NewJsonEventRepository("data.json")

	if err != nil {
		print(err.Error())
		return
	}

	eventsHandler := httpHandler.NewEventsHandler(
		usecase.NewListEventsUseCase(repo),
		usecase.NewGetEventUseCase(repo),
		usecase.NewListSpotsUseCase(repo),
		usecase.NewReserveSpotUseCase(repo),
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpotEvent)
	r.HandleFunc("POST /events/{eventID}/reserve", eventsHandler.ReserveSpot)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		log.Println("Interrupt signal received, starting the graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error on graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Iniciando o servidor HTTP
	log.Println("Server HTTP init on port 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error init server HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Server HTTP finished")
}
