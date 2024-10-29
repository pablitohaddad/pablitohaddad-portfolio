package controllers

import (
	"net/http"
	"os"
	"pablitohaddad-portfolio/database"
	"pablitohaddad-portfolio/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func SendContactMessage(c *gin.Context) {
    var contato models.Contato
    var user models.Usuario

    username := c.Param("username")

    // Buscar o usuário pelo nome de usuário
    if err := database.DatabaseConnection.Where("Username = ?", username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado!"})
        return
    }

    // Atribuindo o id ao meu contato, para deixar dinamico
    contato.UsuarioId = user.ID

    if err := c.ShouldBindJSON(&contato); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    } 

    // Salvando no banco de dados
    database.DatabaseConnection.Create(&contato)

    // Configuração do servidor de e-mail
    d := gomail.NewDialer("smtp.gmail.com", 587, "pablohaddad73@gmail.com", os.Getenv("GOOGLE_API_KEY"))

    // Criando o e-mail que sera enviado para mim
    mailForMe := gomail.NewMessage()
    mailForMe.SetHeader("From", "pablohaddad73@gmail.com")
    mailForMe.SetHeader("To", user.Email)
    mailForMe.SetHeader("Subject", "Nova mensagem de contato de " + contato.NomeRemetente)
    mailForMe.SetBody("text/plain", "Olá "+user.Nome+"!\n" + "Você recebeu uma nova mensagem de contato!\n" +
                            "Nome do remetente: " + contato.NomeRemetente+"\n" +
                            "Email do remetente " + contato.EmailRemetente+"\n" +
                            "Mensagem do remetente: " + contato.Menssagem+"\n"+
                            "Por favor, entre em contato com o email enviado para responde-lo(la)!")

    // Enviando o e-mail
    if err := d.DialAndSend(mailForMe); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao enviar o e-mail"})
        return
    }

    mailForSender := gomail.NewMessage()
    mailForSender.SetHeader("From", "pablohaddad73@gmail.com")
    mailForSender.SetHeader("To", contato.EmailRemetente)
    mailForSender.SetHeader("Subject", "Confirmação de Email - " + user.Nome)
    mailForSender.SetBody("text/plain", "Olá " + contato.NomeRemetente+"! " + "Confirmo que recebi o seu email, irei retornar assim que possível!\n" + 
                            "Sua mensagem:\n" + contato.Menssagem)

    // Enviando o e-mail
    if err := d.DialAndSend(mailForSender); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao enviar o e-mail"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Mensagem enviada com sucesso!"})
}