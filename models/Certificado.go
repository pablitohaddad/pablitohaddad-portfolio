package models

import "gorm.io/gorm"

type Certificado struct{
	gorm.Model
	Titulo string `gorm:"size:255;not null"`
	Instituicao string `gorm:"size:255"`
	DataObtencao string `gorm:"not null"`
	UserID uint // Chave estrangeira pra Usuario
}