package handlers

import (
	"github.com/takama/router"
	"net/http"
)

type Repository struct {
	Username string
	Name     string
}

func CreateRepository(c *router.Control) {
	c.Code(http.StatusCreated)
}

func DeleteRepository(c *router.Control) {
	data := Repository{
		Name: c.Get(":id"),
	}
	c.Code(http.StatusOK).Body(data)
}
