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
    mailForMe.SetHeader("From", "hiscbruh@gmail.com")
    mailForMe.SetHeader("To", "pablohaddad73@gmail.com")
    mailForMe.SetHeader("Subject", "Nova mensagem de contato de " + contato.EmailRemetente)
    mailForMe.SetBody("text/plain", "Nome: "+contato.NomeRemetente+"\nEmail: "+contato.EmailRemetente+"\nMensagem: "+contato.Menssagem)

    // Enviando o e-mail
    if err := d.DialAndSend(mailForMe); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao enviar o e-mail"})
        return
    }

    mailForSender := gomail.NewMessage()
    mailForSender.SetHeader("From", "pablohaddad73@gmail.com")
    mailForSender.SetHeader("To", contato.EmailRemetente)
    mailForSender.SetHeader("Subject", "Confirmação de Email - Pablo Haddad")
    mailForSender.SetBody("text/plain", "Olá " + contato.NomeRemetente+"! " + "Confirmo que recebi o seu email, irei retornar assim que possível!\n" + 
                            "Sua mensagem:\n" + contato.Menssagem)

    // Enviando o e-mail
    if err := d.DialAndSend(mailForSender); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao enviar o e-mail"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Mensagem enviada com sucesso!"})
}
