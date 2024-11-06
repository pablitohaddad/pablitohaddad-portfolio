package controllers

import (
	"net/http"
	"pablitohaddad-portfolio/database"
	"pablitohaddad-portfolio/exceptions"
	"pablitohaddad-portfolio/models"

	"github.com/gin-gonic/gin"
)

// CreateUser lida com a requisição HTTP para criar um novo usuário.
// @Summary Cria um novo usuário
// @Description Recebe dados JSON para criar um novo usuário no sistema
// @Tags Usuários
// @Accept  json
// @Produce  json
// @Param   user body models.Usuario true "Dados do Usuário"
// @Success 201 {object} models.Usuario
// @Failure 400 {object} exceptions.ErrorResponse "Erro ao criar o Usuário"
// @Router /users [post]
func CreateUser(c *gin.Context){
	var user models.Usuario
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, exceptions.ErrorResponse{Error: err.Error()})
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