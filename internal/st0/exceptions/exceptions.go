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

func StorageFull(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "Faild to upload file due to full storage"})
}
