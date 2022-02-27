package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jayndu/stripe-payments/cart"
	"github.com/jayndu/stripe-payments/payments"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	paymentsClientManager := payments.NewClientManager()

	r.Use(CORSMiddleware())

	cart.Setup(paymentsClientManager)
	payments.Setup(r, paymentsClientManager)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
