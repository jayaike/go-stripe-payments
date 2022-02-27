package payments

import (
	"errors"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Service interface {
	initialiseStripePayment(string, string) (string, error)
}

type service struct {
	clientManager ClientManager
}

func newService(cm ClientManager) Service {
	return &service{
		clientManager: cm,
	}
}

func (srvc *service) initialiseStripePayment(clientSlug, itemId string) (string, error) {
	client := srvc.clientManager.findClientBySlug(clientSlug)

	if client == nil {
		return "", errors.New("this service does not exist")
	}

	item, err := client.TXManager.GetItem(itemId)

	if err != nil {
		return "", errors.New("this product does not exist")
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(item.Price)),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	params.AddMetadata("item_id", itemId)

	intent, err := paymentintent.New(params)

	if err != nil {
		return "", errors.New("could not initialise payment. An unexpected error occured")
	}

	return intent.ClientSecret, nil
}
