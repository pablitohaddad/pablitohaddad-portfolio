package models

import "gorm.io/gorm"

type Contato struct {
	gorm.Model
	NomeRemetente string `gorm:"size:255;not null"`
	EmailRemetente string `gorm:"size:255;not null"`
	Menssagem string `gorm:"size:1000;not null"`
	UserID uint // Chave estrangeira pra Usuario
}