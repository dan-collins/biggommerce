package order

import (
	"fmt"
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

//NewClient will create a new order client wrapper based on BC connection details
func NewClient(authToken, authClient, storeKey string) *Client {
	bcClient := connect.NewClient(authToken, authClient, storeKey)
	orderClient := Client{}
	orderClient.BCClient = *bcClient
	orderClient.Limit = 250
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

// GetOrderCount will return an OrderCount struct containing statuses and counts
func (s *Client) GetOrderCount() (*OrderCount, error) {
	var data OrderCount
	err := s.GetAndUnmarshal("v2/orders/count", &data)
	if err != nil {
		return nil, err
	}

	sort.Slice(data.StatusCounts, func(i, j int) bool {
		return data.StatusCounts[i].SortOrder < data.StatusCounts[j].SortOrder
	})
	return &data, nil
}

// GetShipment will return a slice of Shipment structs containing the shipment information
func (s *Client) GetShipment(orderID int) (*[]Shipment, error) {
	url := fmt.Sprintf("v2/orders/%d/shipments", orderID)
	var data []Shipment
	err := s.GetAndUnmarshal(url, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetShipmentsForOrders - Will attempt to concurrently fill the order slice elements with their respective shipment objects from the BC api
func (s *Client) GetShipmentsForOrders(os []Order) (err error) {
	var eg errgroup.Group
	sem := make(chan bool, 20)
	for i := range os {
		j := i
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			shipments, err := s.GetShipment(int(os[j].ID))
			if err != nil {
				return err
			}
			os[j].Shipments = *shipments
			return nil
		})
	}
	err = eg.Wait()
	return
}

// GetShipments will get shipments from the orders returned by the query
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
			var data []Shipment
			url := fmt.Sprintf("v2/orders/%d/shipments", o.ID)
			err := s.GetAndUnmarshal(url, &data)
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
	if q.StatusIDIsZero {
		v.Add("status_id", "0")
	}
	raw = v.Encode()
	return
}

// GetOrderQuery will return an ordered by ID slice of Order structs based on passed in query object
func (s *Client) GetOrderFromRawQuery(rawQuery string) (*[]Order, error) {
	var data []Order
	err := s.GetAndUnmarshalWithQuery("v2/orders/", rawQuery, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrderQuery will return an ordered by ID slice of Order structs based on passed in query object
func (s *Client) GetOrderQuery(oq Query) (*[]Order, error) {
	if oq.Limit == 0 {
		oq.Limit = s.Limit
	}
	var allOrders []Order
	orderCount := s.Limit + 1
	getAllPages := true
	page := 0

	for getAllPages && orderCount >= s.Limit {
		if page == 0 && oq.Page != 0 {
			getAllPages = false
		} else {
			page = page + 1
			oq.Page = page
		}
		rawQuery, err := oq.GetRawQuery()
		if err != nil {
			return nil, err
		}
		data, err := s.GetOrderFromRawQuery(rawQuery)
		if err != nil {
			return nil, err
		}
		orderCount = len(*data)
		allOrders = append(allOrders, *data...)
	}
	sort.Slice(allOrders, func(i, j int) bool {
		return allOrders[i].ID < allOrders[j].ID
	})

	return &allOrders, nil
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
	err = s.GetShipmentsForOrders(*orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrders will return a slice of Order structs based on passed in status
func (s *Client) GetOrders(status int) (*[]Order, error) {
	return s.GetOrderQuery(Query{StatusID: status})
}

// GetOrdersAndProducts will return a slice of Order structs with their products based on passed in status
func (s *Client) GetOrdersAndProducts(status int) (*[]Order, error) {
	orders, err := s.GetOrders(status)
	if err != nil {
		return nil, err
	}
	err = s.GetProductDetail(*orders)

	return orders, err
}

// GetHydratedOrderByID - return a single order with Products, Shipping Addresses, and Coupons Populated.
func (s *Client) GetHydratedOrderByID(orderID string) (order Order, err error) {
	err = s.GetAndUnmarshalWithQuery("v2/orders/", orderID, &order)
	if err != nil {
		return
	}
	err = order.ShippingResource.EagerGet(s, &order.ShippingAddresses)
	if err != nil {
		return
	}
	err = order.CouponResource.EagerGet(s, &order.Coupons)
	if err != nil {
		return
	}
	err = order.ProductResource.EagerGet(s, &order.Products)
	if err != nil {
		return
	}
	shipments, err := s.GetShipment(int(order.ID))
	if err != nil {
		return
	}
	order.Shipments = *shipments
	return
}

// GetAvailableStatuses will return a sorted slice of order statuses from the BC API
func (s *Client) GetAvailableStatuses() (statuses Statuses, err error) {
	err = s.GetAndUnmarshal(
		"v2/order_statuses",
		&statuses,
	)
	if err != nil {
		return
	}

	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].Order < statuses[j].Order
	})
	return
}
