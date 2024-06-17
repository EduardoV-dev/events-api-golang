package api

import (
	"events/internal/api/auth"
	"events/internal/api/event"
	"events/internal/types"
)

func StartAPI(apiServer *types.APIServer) {
	events.RegisterRoutes(apiServer)
	auth.RegisterRoutes(apiServer)
}
