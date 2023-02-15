package helpers

import (
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
	"github.com/gin-gonic/gin"
)

func ToWebResponse(c *gin.Context, data *response.SuccessResponse) {
	c.JSON(data.StatusCode, data)
	return
}
