package handlers

import (
	"log"
)

type Handler struct {
	Stdlog *log.Logger
	Errlog *log.Logger
	Env    map[string]string
}
