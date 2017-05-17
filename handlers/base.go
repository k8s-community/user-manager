package handlers

import (
	"log"
)

// Handler defines
type Handler struct {
	Stdlog *log.Logger
	Errlog *log.Logger
	Env    map[string]string
}
