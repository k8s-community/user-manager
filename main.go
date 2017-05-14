package main

import (
	"github.com/takama/router"
	"github.com/vsaveliev/user-manager/handlers"
)

const (
	apiPrefix = "/api/v1"
)

func Hello(c *router.Control) {
	c.Body("Hello world")
}

func main() {
	r := router.New()
	r.GET("/hello", Hello)

	r.GET(apiPrefix+"/user/:id", handlers.RetrieveUser)
	r.POST(apiPrefix+"/user", handlers.CreateUser)

	r.POST(apiPrefix+"/repository", handlers.CreateRepository)
	r.DELETE(apiPrefix+"/repository/:id", handlers.DeleteRepository)

	r.Listen(":8888")
}
