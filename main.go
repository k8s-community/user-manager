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

// main function
func main() {
	keys := []string{
		"USERMAN_SERVICE_PORT",
		"DOCKER_REGISTRY_SECRET_NAME", "TLS_SECRET_NAME",
		"K8S_HOST", "K8S_TOKEN",
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

	r.PUT(apiPrefix+"/sync-user", h.SyncUser)

	h.Stdlog.Printf("start listening port %s", h.Env["USERMAN_SERVICE_PORT"])

	r.Listen(":" + h.Env["USERMAN_SERVICE_PORT"])
}
