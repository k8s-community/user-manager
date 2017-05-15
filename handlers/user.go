package handlers

import (
	"net/http"

	"github.com/takama/router"
)

type User struct {
	Name string `json:"name"`
}

func (h *Handler) RetrieveUser(c *router.Control) {
	data := User{
		Name: c.Get(":id"),
	}

	c.Code(http.StatusOK).Body(data)
}

func (h *Handler) SyncUser(c *router.Control) {
	/*
		1. namespace с именем как имя аккаунта на github
		2. role и rolebinding в этом namecpace с правами админа в этом namespace
		3. job в Jenkins с указанием на репу с сервисом для которой будет происходить билд
		4. регистрацию еще одной репы из Github в которой находится чарт. Комманда регистрации `helm repo add <user_name>-charts https://github.com/<user_name>/<repo_name>`
	*/

	//createUser(username)
	//createNamespace(namespace)
	//createRole()
	//createRoleBinding()
	//createJenkinsJob(createJobParams)
	//registerChart(username, repName)

	c.Code(http.StatusOK)
}
