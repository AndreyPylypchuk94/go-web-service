package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pylypchuk.home/internal/response"
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

	userId, err := h.authService.ParseToken(token)

	if err != nil {
		response.FailedStopResponse(c, http.StatusUnauthorized, "bad auth data")
		return
	}

	c.Set(userCtx, userId)

	c.Next()
}
