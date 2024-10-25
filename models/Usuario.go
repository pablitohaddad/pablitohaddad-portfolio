package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome string `gorm:"size:255;not null"`
	UrlFotoUsuario string 
	Email string `gorm:"size:255;not null"`
	Senha string `gorm:"size:255;not null"`
	Bio string  `gorm:"size:500"`
}