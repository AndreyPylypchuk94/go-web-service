package api

import (
	"github.com/gin-gonic/gin"
	"pylypchuk.home/internal/service"
)

type Handler struct {
	userService *service.UserWebService
	authService *service.AuthWebService
}

const userCtx = "userCtx"

func NewHandler(userService *service.UserWebService, authService *service.AuthWebService) *Handler {
	return &Handler{userService: userService, authService: authService}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.auth)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getAll)
			users.GET("/:id", h.get)
			users.PUT("/:id", h.update)
			users.DELETE("/:id", h.delete)
		}
	}

	return router
}