package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}
