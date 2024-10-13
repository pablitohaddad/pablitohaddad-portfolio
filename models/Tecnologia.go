package models

import "gorm.io/gorm"

type Tecnologia struct{
	gorm.Model
	Nome string `gorm:"size:255;not null"`
	Projetos []Projeto `gorm:"many2many:project_technologies;"`
}