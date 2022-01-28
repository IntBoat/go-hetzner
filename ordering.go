package hetzner

import (
	"fmt"
	"net/http"
)

// Api: https://robot.your-server.de/doc/webservice/en.html#server-ordering

// OrderingService represents a service to work with server ordering.
// To use the Robot webservice for ordering servers, please activate this function in your Robot administrative
// interface first via "Administration; Settings; Web Service Settings; Ordering".
type OrderingService interface {
	// ListProducts Product overview of currently offered standard server products
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product
	ListProducts() ([]*Product, *http.Response, error)
	// GetProduct Query a specific server product
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product-product-id
	GetProduct(id string) (*Product, *http.Response, error)

	// ListTransactions Overview of all server orders within the last 30 days
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product-product-id
	ListTransactions() ([]*Transaction, *http.Response, error)
	// CreateTransaction Order a new server. If the order is successful, the status code 201 CREATED is returned.
	// See: https://robot.your-server.de/doc/webservice/en.html#post-order-server-transaction
	CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error)
	// GetTransaction Query a specific order transaction
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-transaction-id
	GetTransaction(id string) (*Transaction, *http.Response, error)

	// ListMarketProducts Product overview of currently offered server market products
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product
	ListMarketProducts() ([]*MarketProduct, *http.Response, error)
	// GetMarketProduct Query a specific server market product
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product-product-id
	GetMarketProduct(id string) (*MarketProduct, *http.Response, error)

	// ListMarketTransactions Overview of all server orders within the last 30 days
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product-product-id
	ListMarketTransactions() ([]*MarketTransaction, *http.Response, error)
	// CreateMarketTransaction Order a new server from the server market. If the order is successful, the status code
	// 201 CREATED is returned.
	// See: https://robot.your-server.de/doc/webservice/en.html#post-order-server_market-transaction
	CreateMarketTransaction(req *CreateTransactionRequest) (*MarketTransaction, *http.Response, error)
	// GetMarketTransaction Query a specific order transaction
	// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-transaction-id
	GetMarketTransaction(id string) (*MarketTransaction, *http.Response, error)
}

type OrderingServiceImpl struct {
	client *Client
}

var _ OrderingService = &OrderingServiceImpl{}

func (s *OrderingServiceImpl) ListProducts() ([]*Product, *http.Response, error) {
	path := "/order/server/product"

	data := make([]dataProduct, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Product, len(data))
	for i, d := range data {
		a[i] = d.Product
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) GetProduct(id string) (*Product, *http.Response, error) {
	path := fmt.Sprintf("/order/server/product/%s", id)

	data := dataProduct{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Product, resp, err
}

func (s *OrderingServiceImpl) ListTransactions() ([]*Transaction, *http.Response, error) {
	path := "/order/server/transaction"

	data := make([]dataTransaction, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Transaction, len(data))
	for i, d := range data {
		a[i] = d.Transaction
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error) {
	path := "/order/server/transaction"

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Transaction, resp, err
}

func (s *OrderingServiceImpl) GetTransaction(id string) (*Transaction, *http.Response, error) {
	path := fmt.Sprintf("/order/server/transaction/%v", id)

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Transaction, resp, err
}

func (s *OrderingServiceImpl) ListMarketProducts() ([]*MarketProduct, *http.Response, error) {
	path := "/order/server_market/product"

	data := make([]dataMarketProduct, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*MarketProduct, len(data))
	for i, d := range data {
		a[i] = d.Product
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) GetMarketProduct(id string) (*MarketProduct, *http.Response, error) {
	path := fmt.Sprintf("/order/server_market/product/%s", id)

	data := dataMarketProduct{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Product, resp, err
}

func (s *OrderingServiceImpl) ListMarketTransactions() ([]*MarketTransaction, *http.Response, error) {
	path := "/order/server_market/transaction"

	data := make([]dataMarketTransaction, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*MarketTransaction, len(data))
	for i, d := range data {
		a[i] = d.Transaction
	}
	return a, resp, err
}

func (s *OrderingServiceImpl) CreateMarketTransaction(req *CreateTransactionRequest) (
	*MarketTransaction, *http.Response, error,
) {
	path := "/order/server_market/transaction"

	data := dataMarketTransaction{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Transaction, resp, err
}

func (s *OrderingServiceImpl) GetMarketTransaction(id string) (*MarketTransaction, *http.Response, error) {
	path := fmt.Sprintf("/order/server_market/transaction/%s", id)

	data := dataMarketTransaction{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Transaction, resp, err
}
