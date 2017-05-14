package handlers

import (
	"net/http"

	"github.com/takama/router"
)

type Repository struct {
	Username string
	Name     string
}

func (h *Handler) CreateRepository(c *router.Control) {
	c.Code(http.StatusCreated)
}

func (h *Handler) DeleteRepository(c *router.Control) {
	data := Repository{
		Name: c.Get(":id"),
	}

	c.Code(http.StatusOK).Body(data)
}
