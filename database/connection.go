package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variavel global que chama a conexão do banco de dados
var DatabaseConnection *gorm.DB

// Função que conecta no banco de dados, utilizei o GORM para simplificar, e se der um erro, ele estoura a aplicação
func Connect(){

	// Conecta ao meu banco de dados local Postgres, ele ja tem que estar criado!
	dataSourceName := "host=localhost user=postgres password=admin dbname=meu_portfolio_db port=5432 sslmode=disable"

	var err error
	// GORM abre a conexão com o banco de dados
	DatabaseConnection, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados")
	}

}