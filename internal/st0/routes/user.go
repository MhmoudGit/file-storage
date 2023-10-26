package routes

import (
	"net/http"

	config "github.com/MhmoudGit/file-storage/internal/st0/configs"
	ex "github.com/MhmoudGit/file-storage/internal/st0/exceptions"
	"github.com/MhmoudGit/file-storage/internal/st0/models"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.GET(":id", getUser)
	users.POST("", createUser)
	users.DELETE(":id", deleteUser)

}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	result := config.Db.First(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func createUser(c *gin.Context) {
	var user models.User
	// Bind the request data to the user struct
	if err := c.Bind(&user); err != nil {
		ex.BadRequest(c, err)
		return
	}
	newUser, err := models.NewUser(user.Username)
	if err != nil {
		ex.BadRequest(c, err)
		return
	}
	result := config.Db.Create(newUser)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func deleteUser(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	err := user.DeleteSpace()
	if err != nil {
		ex.BadRequest(c, err)
		return
	}
	result := config.Db.Delete(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
