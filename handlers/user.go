package handlers

import (
	"net/http"

	"github.com/takama/router"
)

type User struct {
	Name string `json:"name"`
}

func RetrieveUser(c *router.Control) {
	data := User{
		Name: c.Get(":id"),
	}
	c.Code(http.StatusOK).Body(data)
}

func CreateUser(c *router.Control) {
	c.Code(http.StatusCreated)
}
