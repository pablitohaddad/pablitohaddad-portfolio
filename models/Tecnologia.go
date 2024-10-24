package models

import "gorm.io/gorm"

type Tecnologia struct{
	gorm.Model
	UrlFoto string `gorm:"not null"` // posso chamar a api devicons para pegar
	Projetos []Projeto `gorm:"many2many:project_technologies;"`
}