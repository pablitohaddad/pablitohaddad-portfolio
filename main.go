package main

import (
	"log"
	"pablitohaddad-portfolio/database"
	"pablitohaddad-portfolio/models"

	"github.com/gin-gonic/gin"
)

func main(){

	// Conecta no meu banco de dados
	database.Connect()

	// Migração automatica de criação de tabelas
	err := database.DatabaseConnection.AutoMigrate(&models.Usuario{}, &models.Projeto{}, &models.Certificado{}, &models.Contato{}, &models.Tecnologia{})
    if err != nil {
        log.Fatal("Falha ao migrar as tabelas:", err)
    }

	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Projeto inicializado com sucesso",
		})
	})

	router.Run(":8080")

}