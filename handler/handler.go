package handler

import (
	"github.com/gin-gonic/gin"
)

// Email is the struct that defines the user email message
type Email struct {
	Name    string `form:"name" json:"name"`
	From    string `form:"from" json:"from"`
	Subject string `form:"subject" json:"subject"`
	Body    string `form:"body" json:"body"`
}

// Response expected from service
type Response struct {
	Status  string `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}

// GetRouter returns a router with the registered endpoints
func GetRouter(s EmailStore) *gin.Engine {
	r := gin.Default()
	// sesWorker := NewSESWorker(sesRegion, sesAccessKey, sesSecretKey)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
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
			Message: "success",
		}
		err := c.ShouldBindJSON(&email)
		if err != nil {
			response = Response{
				Status:  "error",
				Message: err.Error(),
			}
			c.JSON(400, response)
			return
		}
		c.JSON(200, response)
		go s.SendEmail(&email)
		return
	}
	return gin.HandlerFunc(fn)
}
