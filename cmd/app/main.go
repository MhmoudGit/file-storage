package main

import (
	config "github.com/MhmoudGit/file-storage/internal/st0/configs"
	routes "github.com/MhmoudGit/file-storage/internal/st0/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// database connection
	config.Connect()
	config.AutoMigrateDb()

	// define server
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// define routes
	routes.UsersRoutes(r)

	// run server
	r.Run("localhost:8000")
}
