package cart

import "github.com/jayndu/stripe-payments/payments"

type CartTransactionManager struct {
}

func (m *CartTransactionManager) GetItem(id string) (*payments.Item, error) {
	// Typically get this data from a database using the provided id
	// Example: Select cart from cart_table where id = ${id}.
	// We would then calculate the total cost of the cart...
	// That would be the

	return &payments.Item{
		Name:        "Item 1234",
		Description: "Test Description",
		Price:       12394,
	}, nil
}

func (m *CartTransactionManager) CompletePurchase(details *payments.PurchaseDetails) (bool, error) {
	// This method will fulfill the order.
	// Example:
	return true, nil
}
