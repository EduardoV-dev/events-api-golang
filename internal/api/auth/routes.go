package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.RouterGroup) {
  router := server.Group("/auth")
  
  router.POST("/signup", handleSignup)
}
