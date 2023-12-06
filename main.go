package main

import (
	"os"
	"os/signal"
	"polygon-token-oracle-backend/internal/log"
	"polygon-token-oracle-backend/logic/blockchain/events"
	"polygon-token-oracle-backend/logic/config"
	"syscall"
)

func main() {
	_, err := events.NewEvents(config.Get().Event)
	if err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.GetLogger().Info("gracefully shutdown...")
}
