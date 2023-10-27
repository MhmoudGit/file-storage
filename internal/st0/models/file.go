package models

import (
	"os"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Size      int    `gorm:"not null" json:"size"`
	Path      string `gorm:"not null" json:"path"`
	Type      string `gorm:"not null" json:"type"`
	StorageID int    `gorm:"not null" json:"storageID"`
}

func NewFile(name, path, filetype string, size, storageID int) *File {
	return &File{
		Name:      name,
		Size:      size,
		Path:      path,
		Type:      filetype,
		StorageID: storageID,
	}
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
