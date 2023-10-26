package main

import (
	config "github.com/MhmoudGit/file-storage/internal/st0/configs"
)

func main() {
	// database connection
	config.Connect()
	config.AutoMigrateDb()
}
