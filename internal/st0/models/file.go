package models

import (
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
