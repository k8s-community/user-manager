package main

import (
	"log"
	"os"

	"github.com/takama/router"
	"github.com/vsaveliev/user-manager/handlers"
)

const (
	apiPrefix = "/api/v1"
)

func main() {
	keys := []string{
		"USERMAN_SERVICE_PORT",
	}
	h := &handlers.Handler{
		Stdlog: log.New(os.Stdout, "[USERMAN:INFO]: ", log.LstdFlags),
		Errlog: log.New(os.Stderr, "[USERMAN:ERROR]: ", log.LstdFlags),
		Env:    make(map[string]string, len(keys)),
	}
	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			h.Errlog.Fatalf("%s environment variable was not set", key)
		}
		h.Env[key] = value
	}

	r := router.New()

	r.GET(apiPrefix+"/user/:id", h.RetrieveUser)

	r.PUT(apiPrefix+"/sync-user", h.SyncUser)

	r.Listen(":" + h.Env["USERMAN_SERVICE_PORT"])
}
