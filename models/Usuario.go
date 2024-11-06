package models

import (
	"time"
)

type Usuario struct {
	ID uint `json:"id" gorm:"primary_key"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    // DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Nome string `gorm:"size:255;not null"`
	UrlFotoUsuario string `gorm:"not null"`
	Email string `gorm:"size:255;unique;not null"`
	Username string `gorm:"size:255;unique"`
	Senha string `gorm:"size:255;not null"` // armazenar a senha com hash
	Bio string `gorm:"size:500"`
}
