package routes

import (
	"net/http"
	"path/filepath"

	config "github.com/MhmoudGit/file-storage/internal/st0/configs"
	ex "github.com/MhmoudGit/file-storage/internal/st0/exceptions"
	"github.com/MhmoudGit/file-storage/internal/st0/models"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.GET(":userID", getUser)
	users.POST("", createUser)
	users.DELETE(":userID", deleteUser)

	// files
	users.POST(":userID/files", createUserFiles)
	users.DELETE(":userID/files/:fileID", deleteUserFiles)
}

// users controllers
func getUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	result := config.Db.Preload("UserSpace").First(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	user.StorageSize = user.StorageSize / (1024 * 1024)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func createUser(c *gin.Context) {
	var user models.User
	// Bind the request data to the user struct
	if err := c.ShouldBind(&user); err != nil {
		ex.BadRequest(c, err)
		return
	}
	newUser, err := models.NewUser(user.Username, user.Password)
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
	userID := c.Param("userID")
	var user models.User
	data := config.Db.Preload("UserSpace").First(&user, userID)
	if data.Error != nil {
		ex.BadRequest(c, data.Error)
		return
	}
	err := user.DeleteSpace(user.UserSpace.Path)
	if err != nil {
		ex.BadRequest(c, err)
		return
	}
	result := config.Db.Unscoped().Delete(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

// files controllers
func createUserFiles(c *gin.Context) {
	userID := c.Param("userID")

	// get single file from formdata
	file, err := c.FormFile("file")
	if err != nil {
		ex.BadRequest(c, err)
		return
	}

	// get user
	var user models.User
	result := config.Db.Preload("UserSpace").First(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}

	if user.StorageSize <= 0 {
		ex.StorageFull(c)
		return
	}

	// upload file
	destPath := filepath.Join(user.UserSpace.Path, file.Filename)
	fileExists := models.FileExists(destPath)
	if fileExists {
		ex.FileExists(c)
		return
	}

	err = c.SaveUploadedFile(file, destPath)
	if err != nil {
		ex.BadRequest(c, err)
		return
	}

	user.StorageSize -= int(file.Size)
	// Save the updated user object
	update := config.Db.Save(&user)
	if update.Error != nil {
		ex.BadRequest(c, update.Error)
		return
	}

	// save file data to db
	newFile := models.NewFile(file.Filename, user.UserSpace.Path+file.Filename, file.Header.Get("Content-Type"), int(file.Size), user.UserSpace.ID)
	result = config.Db.Create(newFile)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

// files controllers
func deleteUserFiles(c *gin.Context) {
	userID := c.Param("userID")
	fileID := c.Param("fileID")

	// get user
	var user models.User
	result := config.Db.Preload("UserSpace").First(&user, userID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}

	// get file from db
	var file models.File
	result = config.Db.First(&file, fileID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}

	// delete file from server
	err := user.DeleteFile(file.Name)
	if err != nil {
		ex.BadRequest(c, err)
		return
	}

	user.StorageSize += int(file.Size)
	// Save the updated user object
	update := config.Db.Save(&user)
	if update.Error != nil {
		ex.BadRequest(c, update.Error)
		return
	}

	// delet file data to db
	result = config.Db.Unscoped().Delete(&file, fileID)
	if result.Error != nil {
		ex.BadRequest(c, result.Error)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}
