package models

import (
    "gorm.io/gorm"
    "gorm.io/datatypes"

    "github.com/google/uuid"
)

type Product struct {
    gorm.Model
    ID              uint            `gorm:"primaryKey;autoIncrement" json:"ID"`
    Name            string          `json:"name"`
    Description     string          `gorm:"size:255" json:"description"`
    Price           string          `json:"price"`
    CategoryName    string          `gorm:"-" json:"category_name"`
    CategoryID      uint            `gorm:"foreignKey:CategoryID" json:"category_id"`
    ImageURLs       datatypes.JSON  `json:"image_urls"`
}

type Image struct {
    gorm.Model
    ID        uuid.UUID `gorm:"primary_key"`
    Name      string    `gorm:"not null"`
    Type      string    `gorm:"not null"`
    Size      int64     `gorm:"not null"`
    Link      string    `gorm:"not null"`
}

type Category struct {
    gorm.Model
    ID          uint        `gorm.Model:"primaryKey;autoIncrement" json:"ID"`
    Name        string      `gorm:"uniqueIndex;size:255" json:"name"`
    Description string      `gorm:"size:255" json:"description"`
    Products    []Product   `gorm:"foreignKey:CategoryID;references:ID" json:"product"`
}

