package auth

import (
	"events/internal/api/user"
	"events/internal/types"
)

func RegisterRoutes(apiServer *types.APIServer) {
	router := apiServer.APIRouter.Group("/auth")

	userRepo := user.NewUserRepository(apiServer.DB)
	authService := newService(userRepo)
	userService := user.NewUserServices(userRepo)
  
	handler := newHandler(userService, authService)

	router.POST("/signup", handler.signup)
	router.POST("/login", handler.login)
}
