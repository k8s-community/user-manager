package handlers

import (
	"github.com/takama/router"
	"net/http"
)

type User struct {
	Username string `json:"username"`
}

func RetrieveUser(c *router.Control) {
	data := User{
		Username: c.Get(":id"),
	}
	c.Code(http.StatusOK).Body(data)
}

func CreateUser(c *router.Control) {
	c.Code(http.StatusCreated)
}
