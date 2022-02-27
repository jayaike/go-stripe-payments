package payments

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
)

func setupStripe() {
	stripe.Key = "Your stripe secret key"
}

func setupRouter(router *gin.Engine, handler Handler) {
	p := router.Group("/payments")
	{
		p.POST(":service/initialise/", handler.initialisePayment)
	}
}

func Setup(router *gin.Engine, paymentsClientManager ClientManager) {
	service := newService(paymentsClientManager)
	handler := newHandler(service)

	setupStripe()
	setupRouter(router, handler)
}
