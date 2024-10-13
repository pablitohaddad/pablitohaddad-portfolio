package models

import "gorm.io/gorm"

type Projeto struct {
	gorm.Model
	Nome string `gorm: "size:255; not null"`
	Descricao string `gorm: "size:500"`
	LinkReferencial string `gorm: "size:255`
	UsuarioId uint // Chave Estrangeira pra Usuario
	Tecnologias []Tecnologia `gorm:"many2many:project_technologies;"` // a relação muitos para muitos 

}