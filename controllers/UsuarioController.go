package controllers

import (
	"net/http"
	"pablitohaddad-portfolio/database"
	"pablitohaddad-portfolio/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){ // c é a requisição HTTP que recebemos e vamos receber
	var user models.Usuario

	// Pegar os dados da requisição e vincula ao nosso usuário com ShouldBindJSON.
	// err tem dois estados:
	// 1. Valor de erro que vai ser tratado dentro do if
	// 2. nil, que indica que não houve erro ao fazer ao chamar a função ShouldBindJSON
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DatabaseConnection.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context){
	var users []models.Usuario
	database.DatabaseConnection.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context){
	var user models.Usuario
	id := c.Param("id")
	if err := database.DatabaseConnection.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
        return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    var user models.Usuario
    id := c.Param("id")

	// verifico se o usuario existe no banco de dados
    if err := database.DatabaseConnection.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
        return
    }
	// vinculo os dados da requisição com os meus existentes
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	// Atualizando o meu usuario
	database.DatabaseConnection.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context){
	var user models.Usuario
	id := c.Param("id")
	if err := database.DatabaseConnection.First(&user, id).Error; err!= nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	database.DatabaseConnection.Delete(&user)
	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted sucessfully"})
}


