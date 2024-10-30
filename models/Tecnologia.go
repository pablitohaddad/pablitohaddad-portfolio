package models

import "gorm.io/gorm"

type Tecnologia struct {
	gorm.Model
	Nome      string     `gorm:"size:100;not null;unique"` // Nome da tecnologia
	UrlFoto   string     `gorm:"not null"` // Link para o ícone da tecnologia
	Projetos  []Projeto  `gorm:"many2many:project_technologies;"`
}