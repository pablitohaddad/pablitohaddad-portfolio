package models

import "gorm.io/gorm"

type Certificado struct{
	gorm.Model
	Titulo string `gorm:"size:255;not null"`
	Instituicao string `gorm:"size:255"` // Posso chamar a api de instituições de ensino
	DataObtencao string `gorm:"not null"`
	CargaHoraria string `gorm:"not null"`
	UrlCertificado string `gorm:"not null"`
	UserID uint // Chave estrangeira pra Usuario
}