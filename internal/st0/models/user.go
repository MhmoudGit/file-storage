package models

import (
	"io"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

const Path = "./storage/"

type User struct {
	gorm.Model
	Username    string  `gorm:"not null" json:"username" form:"username"`
	Password    string  `gorm:"not null" json:"-" form:"password"`
	StorageSize int     `gorm:"not null" json:"storageSize"`
	UserSpace   Storage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"-"`
}

func NewUser(username string) (*User, error) {
	storageSizeInKb := 1024 * 1024 * 100
	spaceName := username + "-space"
	err := os.Mkdir(Path+spaceName, os.ModeDir)
	if err != nil {
		return &User{}, err
	}
	return &User{
		Username:    username,
		StorageSize: storageSizeInKb,
		UserSpace: Storage{
			Name: spaceName,
			Path: Path + spaceName + "/",
		},
	}, nil
}

func (u *User) UploadFile(file *os.File) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	u.StorageSize -= int(fileInfo.Size())

	if u.StorageSize <= 0 {
		return err
	}
	filePath := filepath.Join(u.UserSpace.Path, file.Name())
	filePtr, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer filePtr.Close()
	// Copy the uploaded file data to the new file
	_, err = io.Copy(filePtr, file)
	if err != nil {
		return err
	}
	return nil
}

func (u User) DeleteFile(fileName string) error {
	filePath := filepath.Join(u.UserSpace.Path, fileName)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (u User) DeleteSpace(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
