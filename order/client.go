package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/dan-collins/biggommerce/connect"
	"github.com/google/go-querystring/query"
	"golang.org/x/sync/errgroup"
)

// Client is a wrapper struct that embeds the BCClient from the client package. It handles connection to the BigCommerce API
type Client struct {
	connect.BCClient
}

//NewClient will create a new client wrapper based on BC connection details
func NewClient(authToken, authClient, storeKey string) *Client {
	bcClient := connect.NewClient(authToken, authClient, storeKey)
	orderClient := Client{}
	orderClient.BCClient = *bcClient
	return &orderClient
}

// GetProductDetail - Will attempt to concurrently fill the order slice elements with their respective products from the BC api
func (s *Client) GetProductDetail(os []Order) (err error) {
	var eg errgroup.Group
	sem := make(chan bool, 20)
	for i := range os {
		j := i
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			return os[j].ProductResource.EagerGet(s, &os[j].Products)
		})
	}
	err = eg.Wait()
	return
}

// GetShippingAddressesForOrders - Will attempt to concurrently fill the order slice elements with their respective shipping addresses from the BC api
func (s *Client) GetShippingAddressesForOrders(os []Order) (err error) {
	var eg errgroup.Group
	sem := make(chan bool, 20)
	for i := range os {
		j := i
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			return os[j].ShippingResource.EagerGet(s, &os[j].ShippingAddresses)
		})
	}
	err = eg.Wait()
	return
}

// GetCouponsForOrders - Will attempt to concurrently fill the order slice elements with their respective coupon objects from the BC api
func (s *Client) GetCouponsForOrders(os []Order) (err error) {
	var eg errgroup.Group
	sem := make(chan bool, 20)
	for i := range os {
		j := i
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			return os[j].CouponResource.EagerGet(s, &os[j].Coupons)
		})
	}
	err = eg.Wait()
	return
}

// GetOrderCount Return an OrderCount struct containing statuses and counts
func (s *Client) GetOrderCount() (*OrderCount, error) {
	url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/count", s.StoreKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var data OrderCount
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	sort.Slice(data.StatusCounts, func(i, j int) bool {
		return data.StatusCounts[i].SortOrder < data.StatusCounts[j].SortOrder
	})
	return &data, nil
}

// GetShipments get shipments from the orders returned by the query
func (s *Client) GetShipments(oq Query) ([]Shipment, error) {
	os, err := s.GetOrderQuery(oq)
	if err != nil {
		return nil, err
	}
	var eg errgroup.Group
	sem := make(chan bool, 20)
	shipChan := make(chan Shipment, 3)
	for _, o := range *os {
		o := o
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/%d/shipments", s.StoreKey, o.ID)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return err
			}
			res, err := s.DoRequest(req)
			if err != nil {
				return err
			}

			var data []Shipment
			err = json.Unmarshal(res, &data)
			if err != nil {
				return err
			}

			for _, ship := range data {
				shipChan <- ship
			}
			return nil
		})
	}

	shipments := make([]Shipment, 0)
	go func() {
		for i := range shipChan {
			shipments = append(shipments, i)
		}
	}()
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	return shipments, nil
}

// GetRawQuery gets the struct in query string form
func (q Query) GetRawQuery() (raw string, err error) {
	if !q.MinDateCreated.IsZero() {
		q.MinDateCreatedRaw = q.MinDateCreated.Format(time.RFC1123Z)
	}
	if !q.MaxDateCreated.IsZero() {
		q.MaxDateCreatedRaw = q.MaxDateCreated.Format(time.RFC1123Z)
	}
	if !q.MinDateModified.IsZero() {
		q.MinDateModifiedRaw = q.MinDateModified.Format(time.RFC1123Z)
	}
	if !q.MaxDateModified.IsZero() {
		q.MaxDateModifiedRaw = q.MaxDateModified.Format(time.RFC1123Z)
	}
	v, err := query.Values(q)
	if err != nil {
		return "", err
	}
	raw = v.Encode()
	return
}

// GetOrderQuery Return an ordered by ID slice of Order structs based on passed in query object
func (s *Client) GetOrderQuery(oq Query) (*[]Order, error) {
	url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/", s.StoreKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery, err = oq.GetRawQuery()
	if err != nil {
		return nil, err
	}

	res, err := s.DoRequest(req)
	if err != nil {
		return nil, err
	}

	var data []Order
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
	return &data, nil
}

// GetHydratedOrders Return a slice of Order structs based on passed in query object
func (s *Client) GetHydratedOrders(oq Query) (*[]Order, error) {
	orders, err := s.GetOrderQuery(oq)
	if err != nil {
		return nil, err
	}
	// hydrate products
	err = s.GetProductDetail(*orders)
	if err != nil {
		return nil, err
	}
	// hydrate addresses
	err = s.GetShippingAddressesForOrders(*orders)
	if err != nil {
		return nil, err
	}
	// hydrate coupons
	err = s.GetCouponsForOrders(*orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrders Return a slice of Order structs based on passed in status
func (s *Client) GetOrders(status int) (*[]Order, error) {
	return s.GetOrderQuery(Query{StatusID: status})
}

// GetOrdersAndProducts Return a slice of Order structs based on passed in status
func (s *Client) GetOrdersAndProducts(status int) (*[]Order, error) {
	orders, err := s.GetOrders(status)
	if err != nil {
		return nil, err
	}
	err = s.GetProductDetail(*orders)

	return orders, err
}

// GetOrderByID - return a single order with Shipping Address Populated.
func (s *Client) GetOrderByID(orderID string) (order Order, err error) {
	url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/%s", s.StoreKey, orderID)

	body, err := s.GetBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return
	}
	err = order.ShippingResource.EagerGet(s, &order.ShippingAddresses)
	if err != nil {
		return
	}
	err = order.ProductResource.EagerGet(s, &order.Products)
	return
}
