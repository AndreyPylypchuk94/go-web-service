package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pylypchuk.home/internal/model"
	"pylypchuk.home/internal/request"
	"pylypchuk.home/internal/response"
)

func (h *Handler) signIn(c *gin.Context) {
	var requestDto request.LoginRequest
	if err := c.BindJSON(&requestDto); err != nil {
		response.FailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.authService.SignIn(requestDto)
	if err != nil {
		response.FailedResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.DataResponse(c, map[string]string{"token": token})
}

func (h *Handler) signUp(c *gin.Context) {
	var dto model.CreateUser

	if err := c.BindJSON(&dto); err != nil {
		response.FailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authService.SignUp(dto); err != nil {
		response.FailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.StatusResponse(c, http.StatusCreated)
}
