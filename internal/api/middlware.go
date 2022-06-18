package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"pylypchuk.home/internal/response"
	"pylypchuk.home/internal/service/userStorage"
	"strings"
)

func (h *Handler) auth(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if "" == header || !strings.HasPrefix(header, "Bearer ") {
		response.FailedStopResponse(c, http.StatusUnauthorized, "auth header not present or incorrect")
		return
	}

	parts := strings.Split(header, " ")

	if len(parts) != 2 {
		response.FailedStopResponse(c, http.StatusUnauthorized, "auth header not present or incorrect")
		return
	}

	token := parts[1]

	userId, roles, err := h.authService.ParseToken(token)

	if err != nil {
		response.FailedStopResponse(c, http.StatusUnauthorized, "bad auth data")
		return
	}

	if !userStorage.Contains(userId) {
		response.FailedStopResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	c.Set(userIdCtx, userId)
	c.Set(userRolesCtx, roles)

	c.Next()
}

func (h *Handler) hasAdminRole(c *gin.Context) {
	roles := c.GetStringSlice(userRolesCtx)

	if !slices.Contains(roles, "admin") {
		response.FailedStopResponse(c, http.StatusForbidden, "no access")
		return
	}

	c.Next()
}
