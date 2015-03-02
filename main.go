package main

import (
	"github.com/sintell/wsgame/server"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	gameServer := server.New()

	err := gameServer.Cfg.Load()
	if err != nil {
		gameServer.Logger.Error("Can't load config, reason: %s", err.Error())
	}

	gameServer.Logger.Debug("Load config file: %i", *gameServer.Cfg.Get())

	http.HandleFunc("/game", gameServer.Handler())
	gameServer.Logger.Info("Starting server at %s", ":8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		logger.Fatalf("Can't start at %s because of: %S", ":8080", err.Error())
		panic(err.Error())
	}
}
