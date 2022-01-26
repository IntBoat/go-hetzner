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
	ListProducts() ([]*Product, *http.Response, error)
	GetProduct(id string) (*Product, *http.Response, error)

	ListTransactions() ([]*Transaction, *http.Response, error)
	CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error)
	GetTransaction(id string) (*Transaction, *http.Response, error)

	ListMarketProducts() ([]*MarketProduct, *http.Response, error)
	GetMarketProduct(id string) (*Product, *http.Response, error)

	ListMarketTransactions() ([]*Transaction, *http.Response, error)
	CreateMarketTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error)
	GetMarketTransaction(id string) (*Transaction, *http.Response, error)
}

type OrderingServiceImpl struct {
	client *Client
}

var _ OrderingService = &OrderingServiceImpl{}

// ListProducts Product overview of currently offered standard server products
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product
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

// GetProduct Query a specific server product
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product-product-id
func (s *OrderingServiceImpl) GetProduct(id string) (*Product, *http.Response, error) {
	path := fmt.Sprintf("/order/server/product/%s", id)

	data := dataProduct{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Product, resp, err
}

// ListTransactions Overview of all server orders within the last 30 days
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product-product-id
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

// CreateTransaction Order a new server. If the order is successful, the status code 201 CREATED is returned.
// See: https://robot.your-server.de/doc/webservice/en.html#post-order-server-transaction
func (s *OrderingServiceImpl) CreateTransaction(req *CreateTransactionRequest) (*Transaction, *http.Response, error) {
	path := "/order/server/transaction"

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Transaction, resp, err
}

// GetTransaction Query a specific order transaction
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-transaction-id
func (s *OrderingServiceImpl) GetTransaction(id string) (*Transaction, *http.Response, error) {
	path := fmt.Sprintf("/order/server/transaction/%v", id)

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Transaction, resp, err
}

// ListMarketProducts Product overview of currently offered server market products
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product
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

// GetMarketProduct Query a specific server market product
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product-product-id
func (s *OrderingServiceImpl) GetMarketProduct(id string) (*Product, *http.Response, error) {
	path := fmt.Sprintf("/order/server_market/product/%s", id)

	data := dataProduct{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Product, resp, err
}

// ListMarketTransactions Overview of all server orders within the last 30 days
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-product-product-id
func (s *OrderingServiceImpl) ListMarketTransactions() ([]*Transaction, *http.Response, error) {
	path := "/order/server_market/transaction"

	data := make([]dataTransaction, 0)
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)

	a := make([]*Transaction, len(data))
	for i, d := range data {
		a[i] = d.Transaction
	}
	return a, resp, err
}

// CreateMarketTransaction Order a new server from the server market. If the order is successful, the status code
// 201 CREATED is returned.
// See: https://robot.your-server.de/doc/webservice/en.html#post-order-server_market-transaction
func (s *OrderingServiceImpl) CreateMarketTransaction(req *CreateTransactionRequest) (
	*Transaction, *http.Response, error,
) {
	path := "/order/server_market/transaction"

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodPost, path, req, &data)
	return data.Transaction, resp, err
}

// GetMarketTransaction Query a specific order transaction
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server_market-transaction-id
func (s *OrderingServiceImpl) GetMarketTransaction(id string) (*Transaction, *http.Response, error) {
	path := fmt.Sprintf("/order/server_market/transaction/%s", id)

	data := dataTransaction{}
	resp, err := s.client.Call(http.MethodGet, path, nil, &data)
	return data.Transaction, resp, err
}
