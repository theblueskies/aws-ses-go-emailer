package handler

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Email is the struct that defines the user email message
type Email struct {
	Name    string `form:"name" json:"name"`
	From    string `form:"from" json:"from"`
	Subject string `form:"subject" json:"subject"`
	Body    string `form:"body" json:"body"`
}

// ComposeText puts together the Name, From and Body into one string to send in the final email
func (e *Email) ComposeText() string {
	return fmt.Sprintf("Name: %s \n \n Email: %s \n \n Message: %s ", e.Name, e.From, e.Body)
}

// Response expected from service
type Response struct {
	Status  string `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}

// DefaultMessage is what we make the subject of the email if there is no subject present
const DefaultMessage = "Message from DPA4u"

// GetRouter returns a router with the registered endpoints
func GetRouter(s EmailStore) *gin.Engine {
	r := gin.Default()
	// config := cors.DefaultConfig()
	r.Use(cors.Default())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "email app online",
		})
	})
	r.POST("/email", HandleEmail(s))
	return r
}

// HandleEmail handles the POST body and uses AWS SES to send the email
func HandleEmail(s EmailStore) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var email Email
		response := Response{
			Status:  "success",
			Message: "email app online",
		}
		err := c.ShouldBind(&email)
		if err != nil {
			response = Response{
				Status:  "error",
				Message: err.Error(),
			}
			c.JSON(400, response)
			return
		}
		c.JSON(200, response)
		if email.Subject == "" {
			email.Subject = DefaultMessage
		}
		// Only send the email if there is a body
		if email.Body != "" {
			go s.SendEmail(&email)
		}
		return
	}
	return gin.HandlerFunc(fn)
}
