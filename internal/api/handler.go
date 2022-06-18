package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pylypchuk.home/internal/service"
	"pylypchuk.home/pkg/context"
)

type Handler struct {
	userService *service.UserWebService
	authService *service.AuthWebService
}

const userIdCtx = "userIdCtx"
const userRolesCtx = "userRolesCtx"

func NewHandler() *Handler {
	return &Handler{
		userService: context.Get("userWebService").(*service.UserWebService),
		authService: context.Get("authWebService").(*service.AuthWebService),
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	root := router.Group("/")
	{
		root.GET("/", h.root)

		auth := root.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn)
			auth.POST("/sign-up", h.signUp)
		}

		api := root.Group("/api", h.auth)
		{
			users := api.Group("/users")
			{
				users.GET("/", h.getAll)
				users.GET("/:id", h.get)
				users.PUT("/:id", h.hasAdminRole, h.update)
				users.DELETE("/:id", h.hasAdminRole, h.delete)
			}
		}
	}

	return router
}

func (h *Handler) root(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello Go")
}
