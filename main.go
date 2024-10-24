package main

import (
	"log"
	"pablitohaddad-portfolio/controllers"
	"pablitohaddad-portfolio/database"
	"pablitohaddad-portfolio/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init() {
    // Carregar o arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func main(){



	// Conecta no meu banco de dados
	database.Connect()

	// Migração automatica de criação de tabelas
	err := database.DatabaseConnection.AutoMigrate(&models.Usuario{}, &models.Projeto{}, &models.Certificado{}, &models.Contato{}, &models.Tecnologia{})
    if err != nil {
        log.Fatal("Falha ao migrar as tabelas:", err)
    }

	router := gin.Default()

	// ROTAS DO USUARIO
	router.POST("/createUser", controllers.CreateUser) // Definindo a rota do create
	router.GET("/getUsers", controllers.GetAllUsers) // Definindo a rota do get all users
	router.GET("/getUser/:id", controllers.GetUserById) // Definindo a rota do get all users
	router.PUT("/updateUser/:id", controllers.UpdateUser) // Definindo a rota do update user
	router.DELETE("/deleteUser/:id", controllers.DeleteUser) // Definindo a rota do update user


	// ROTA DE ENVIAR MENSAGEM
	router.POST("/sendMessageForMe", controllers.SendContactMessage)

	router.Run(":8080")

}