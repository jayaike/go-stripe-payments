package cart

import (
	"fmt"

	"github.com/jayndu/stripe-payments/payments"
)

func setupPayments(clientManager payments.ClientManager) {
	clientManager.RegisterClient(&payments.Client{
		Name:      "cart",
		Slug:      "cart",
		TXManager: &CartTransactionManager{},
	})
}

func Setup(paymentsClientManager payments.ClientManager) {
	fmt.Println("setting up cart...")

	setupPayments(paymentsClientManager)
}
