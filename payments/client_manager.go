package payments

import (
	"log"
)

type TransactionManager interface {
	GetItem(string) (*Item, error)
	CompletePurchase(*PurchaseDetails) (bool, error)
}

type ClientManager interface {
	RegisterClient(*Client)
	findClientBySlug(string) *Client
}

type Item struct {
	Name        string
	Description string
	Price       float32
}

type PurchaseDetails struct {
	ItemID string
}

type Client struct {
	Name      string
	Slug      string
	TXManager TransactionManager
}

type clientManager struct {
	clients map[string]*Client
}

func (m *clientManager) RegisterClient(c *Client) {
	if _, ok := m.clients[c.Slug]; ok {
		log.Fatalf("%v is already a registered service", c.Name)
	}

	m.clients[c.Slug] = c
}

func (m *clientManager) findClientBySlug(clientSlug string) *Client {
	client, ok := m.clients[clientSlug]

	if ok {
		return client
	}

	return nil
}

func NewClientManager() ClientManager {
	return &clientManager{
		clients: make(map[string]*Client),
	}
}
