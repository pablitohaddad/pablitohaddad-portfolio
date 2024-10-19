package main

import (
	"log"
	"pablitohaddad-portfolio/controllers"
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

	router.POST("/createUser", controllers.CreateUser) // Definindo a rota do create
	router.GET("/getUser", controllers.GetAllUsers) // Definindo a rota do get all users
	router.GET("/getUser/:id", controllers.GetUserById) // Definindo a rota do get all users
	router.PUT("/updateUser/:id", controllers.UpdateUser) // Definindo a rota do update user
	router.DELETE("/deleteUser/:id", controllers.DeleteUser) // Definindo a rota do update user

	router.Run(":8080")

}