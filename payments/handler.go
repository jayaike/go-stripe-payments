package payments

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	initialisePayment(*gin.Context)
}

type handler struct {
	service Service
}

func newHandler(s Service) Handler {
	return &handler{
		service: s,
	}
}

func (h *handler) initialisePayment(c *gin.Context) {
	var req InitialisePaymentRequest
	clientSlug := c.Params.ByName("service")

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Could not initialise payment",
			Error:   err.Error(),
		})
		return
	}

	stripeClientSecret, err := h.service.initialiseStripePayment(clientSlug, req.ItemId)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Message: "Failed to get checkout url",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Message: "Payment initialised successfully",
		Data: &InitialisePaymentResponse{
			ClientSecret: stripeClientSecret,
		},
	})
}
