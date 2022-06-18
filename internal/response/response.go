package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"error_message"`
}

func DataResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{
		Data:         data,
		Success:      true,
		ErrorMessage: "",
	})
}

func StatusResponse(c *gin.Context, status int) {
	c.Status(status)
}

func FailedResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, response{
		Data:         nil,
		Success:      false,
		ErrorMessage: errorMessage,
	})
}

func FailedStopResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, response{
		Data:         nil,
		Success:      false,
		ErrorMessage: errorMessage,
	})
}
