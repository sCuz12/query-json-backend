package middleware

import (
	"jql-server/data"

	"github.com/sCuz12/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}