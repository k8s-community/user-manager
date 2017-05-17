package handlers

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/takama/router"
	"github.com/vsaveliev/user-manager/k8s"
)

// User defines
type User struct {
	Name string `json:"name"`
}

// SyncUser activates user in k8s system (creates namespaces, secrets)
func (h *Handler) SyncUser(c *router.Control) {
	var user User

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.Errlog.Printf("couldn't read request body: %s", err)
		c.Code(http.StatusBadRequest).Body(nil)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		h.Errlog.Printf("couldn't validate request body: %s", err)
		c.Code(http.StatusBadRequest).Body(nil)
		return
	}

	if user.Name == "" {
		c.Code(http.StatusBadRequest).Body(nil)
		return
	}

	h.Stdlog.Printf("try to activate user %s", user.Name)

	client, err := k8s.NewClient(h.Env["K8S_HOST"], h.Env["K8S_TOKEN"])
	if err != nil {
		h.Errlog.Printf("cannot connect to k8s server: %s", err)
		c.Code(http.StatusInternalServerError).Body(nil)
		return
	}

	namespace, _ := client.GetNamespace(user.Name)
	if namespace != nil {
		h.Stdlog.Printf("user %s already exists", user.Name)
		c.Code(http.StatusOK).Body(nil)
		return
	}

	err = client.CreateNamespace(user.Name)
	if err != nil {
		h.Errlog.Printf("%s", err)
		c.Code(http.StatusInternalServerError).Body(nil)
		return
	}

	secretNames := []string{h.Env["DOCKER_REGISTRY_SECRET_NAME"], h.Env["TLS_SECRET_NAME"]}
	for _, secretName := range secretNames {
		err = client.CopySecret(secretName, "default", user.Name)
		if err != nil {
			h.Errlog.Printf("%s", err)
			c.Code(http.StatusInternalServerError).Body(nil)
			return
		}
	}

	c.Code(http.StatusOK).Body(nil)

	h.Stdlog.Printf("user %s is activated", user.Name)
}
