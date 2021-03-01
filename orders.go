package biggommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	. "github.com/dan-collins/biggommerce/model"
	"golang.org/x/sync/errgroup"
)

// GetProductDetail - Will attempt to concurrently fill the order slice elements with their respective products from the BC api
func (s *BCClient) GetProductDetail(os []Order) (err error) {
	var eg errgroup.Group
	sem := make(chan bool, 20)
	for i := range os {
		i := i
		eg.Go(func() error {
			sem <- true
			defer func() { <-sem }()
			return os[i].ProductResource.EagerGet(s, &os[i].Products)
		})
	}
	err = eg.Wait()
	return
}

// GetOrderCount Return an OrderCount struct containing statuses and counts
func (s *BCClient) GetOrderCount() (*OrderCount, error) {
	url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/count", s.StoreKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.doRequest(req)
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
func (s *BCClient) GetShipments(oq OrderQuery) ([]Shipment, error) {
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
			res, err := s.doRequest(req)
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

// GetOrderQuery Return a slice of Order structs based on passed in query object
func (s *BCClient) GetOrderQuery(oq OrderQuery) (*[]Order, error) {
	url := fmt.Sprintf(s.BaseURL+"%s/v2/orders/", s.StoreKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery, err = oq.GetRawQuery()
	if err != nil {
		return nil, err
	}

	res, err := s.doRequest(req)
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

// GetOrders Return a slice of Order structs based on passed in status
func (s *BCClient) GetOrders(status int) (*[]Order, error) {
	return s.GetOrderQuery(OrderQuery{StatusID: status})
}

// GetOrdersAndProducts Return a slice of Order structs based on passed in status
func (s *BCClient) GetOrdersAndProducts(status int) (*[]Order, error) {
	orders, err := s.GetOrders(status)
	if err != nil {
		return nil, err
	}
	err = s.GetProductDetail(*orders)

	return orders, err
}

// GetOrderByID - return a single order with Shipping Address Populated.
func (s *BCClient) GetOrderByID(orderID string) (order Order, err error) {
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
