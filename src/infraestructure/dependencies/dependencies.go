package dependencies

import (
	"github.com/adelmar92/GoApiClean/src/entrypoints"
	"github.com/adelmar92/GoApiClean/src/entrypoints/handlers/ping"
)

type HandlerContainer struct {
	Ping entrypoints.Handlers
}

func Start() HandlerContainer {
	//handlers
	pingHandler := ping.NewPing()

	handlers := HandlerContainer{
		Ping: pingHandler,
	}

	return handlers
}
