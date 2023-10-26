package models

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

const Path = "./storage/"

type User struct {
	gorm.Model
	Username    string  `gorm:"not null" json:"username"`
	Password    string  `gorm:"not null" json:"-"`
	StorageSize int     `gorm:"not null" json:"storageSize"`
	UserSpace   Storage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"-"`
}

func CreateUser(username string) User {
	storageSizeInKb := 1024 * 1024 * 100
	spaceName := username + "-space"
	err := os.Mkdir(Path+spaceName, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	return User{
		Username:    username,
		StorageSize: storageSizeInKb,
		UserSpace: Storage{
			Name: spaceName,
			Path: Path + spaceName + "/",
		},
	}
}

func (u *User) UploadFile(file *os.File) {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	u.StorageSize -= int(fileInfo.Size())

	if u.StorageSize <= 0 {
		log.Fatal("Looks like you are out of storage")
		return
	}
	filePath := filepath.Join(u.UserSpace.Path, file.Name())
	filePtr, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
	// Copy the uploaded file data to the new file
	_, err = io.Copy(filePtr, file)
	if err != nil {
		log.Fatal(err)
	}
}

func (u User) DeleteFile(fileName string) {
	filePath := filepath.Join(u.UserSpace.Path, fileName)
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
