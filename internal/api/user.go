package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pylypchuk.home/internal/model"
	"pylypchuk.home/internal/response"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	var dto model.CreateUser

	err := c.BindJSON(&dto)
	if err != nil {
		response.FailedResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}

	id, err := h.userService.Update(dto)
	if err != nil {
		response.FailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.DataResponse(c, id)
}

func (h *Handler) getAll(c *gin.Context) {
	response.DataResponse(c, h.userService.Get())
}

func (h *Handler) get(c *gin.Context) {
	id, err := readIdPathVariable(c, "id")
	if err != nil {
		response.FailedResponse(c, http.StatusNotFound, err.Error())
		return
	}

	user, err := h.userService.GetById(id)

	if err != nil {
		response.FailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.DataResponse(c, user)
}

func (h *Handler) delete(c *gin.Context) {
	id, err := readIdPathVariable(c, "id")
	if err != nil {
		response.FailedResponse(c, http.StatusNotFound, err.Error())
		return
	}

	err = h.userService.DeleteById(id)
	if err != nil {
		response.FailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.StatusResponse(c, http.StatusOK)
}

func readIdPathVariable(c *gin.Context, variable string) (int64, error) {
	idParam := c.Param(variable)
	id, err := strconv.ParseInt(idParam, 0, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
