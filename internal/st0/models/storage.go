package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	Path   string `gorm:"not null" json:"path"`
	Files  []File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:StorageID" json:"-"`
	UserID int    `gorm:"not null" json:"userID"`
}
