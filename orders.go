package biggommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/sync/errgroup"
)

//StatusCount struct representing the individual statuses and count returned by BigCommerce GET /orders/count
type StatusCount struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SystemLabel string `json:"system_label"`
	CustomLabel string `json:"custom_label"`
	SystemDesc  string `json:"system_description"`
	Count       int    `json:"count"`
	SortOrder   int    `json:"sort_order"`
}

// OrderCount struct representing the return body of BigCommerce GET /orders/count
type OrderCount struct {
	StatusCounts []StatusCount `json:"statuses"`
	Count        int           `json:"count"`
}

// Address struct representing a BigCommerce address resource
type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string
	StreetOne string `json:"street_1"`
	StreetTwo string `json:"street_2"`
	City      string
	State     string
	Zip       string
	Country   string
	Phone     string
	Email     string
}

// ShippingAddress struct representing a BigCommerce shipping address resource
type ShippingAddress struct {
	Address
	ShippingMethod string `json:"shipping_method"`
}

// BCDate struct representing a BigCommerce date resource
type BCDate struct {
	time.Time
}

// UnmarshalJSON unmarshall date into time object
func (bcD *BCDate) UnmarshalJSON(input []byte) error {
	strTime := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(time.RFC1123Z, strTime)
	if err != nil {
		return err
	}
	bcD.Time = newTime
	return nil
}

// Shipment struct representing BC shipment
type Shipment struct {
	ID                   int64       `json:"id"`
	OrderID              int64       `json:"order_id"`
	CustomerID           int64       `json:"customer_id"`
	OrderAddressID       int64       `json:"order_address_id"`
	DateCreated          string      `json:"date_created"`
	TrackingNumber       string      `json:"tracking_number"`
	MerchantShippingCost float64     `json:"merchant_shipping_cost,string"`
	ShippingMethod       string      `json:"shipping_method"`
	Comments             string      `json:"comments"`
	ShippingProvider     string      `json:"shipping_provider"`
	TrackingCarrier      string      `json:"tracking_carrier"`
	TrackingLink         string      `json:"tracking_link"`
	BillingAddress       Address     `json:"billing_address"`
	ShippingAddress      Address     `json:"shipping_address"`
	Items                []OrderItem `json:"items"`
}

// OrderItem struct representing BC orderitem
type OrderItem struct {
	OrderProductID int64 `json:"order_product_id"`
	ProductID      int64 `json:"product_id"`
	Quantity       int64 `json:"quantity"`
}

// Order struct representing the return body of BigCommerce GET /orders
type Order struct {
	ID                int      `json:"id"`
	ItemsTotal        int      `json:"items_total"`
	ItemsShipped      int      `json:"items_shipped"`
	StaffNotes        string   `json:"staff_notes"`
	CustomerMessage   string   `json:"customer_message"`
	ProductResource   Resource `json:"products"`
	Products          []OrderProduct
	BillingAddress    Address  `json:"billing_address"`
	ShippingResource  Resource `json:"shipping_addresses"`
	ShippingAddresses []ShippingAddress
	Date              BCDate `json:"date_created"`
	DateShipped       BCDate `json:"date_shipped"`
}

// OrderProduct struct representing the product sub object that is part of the Order
type OrderProduct struct {
	ID               int                  `json:"id"`
	ProductID        int                  `json:"product_id"`
	Name             string               `json:"name"`
	ProductOptions   []OrderProductOption `json:"product_options"`
	Quantity         int                  `json:"quantity"`
	QuantityRefunded int                  `json:"quantity_refunded"`
	SKU              string               `json:"sku"`
}

// OrderProductOption struct representing the product option detail objects that are part of the order product
type OrderProductOption struct {
	DisplayName  string `json:"display_name"`
	DisplayValue string `json:"display_value"`
	Value        string `json:"value"`
	Type         string `json:"type"`
}

// Resource general struct for resource url and type
type Resource struct {
	URL      string
	Resource string
}

// OrderQuery struct to handle orders endpoint search query params
type OrderQuery struct {
	MinID              int       `url:"min_id,omitempty"`
	MaxID              int       `url:"max_id,omitempty"`
	MinTotal           float64   `url:"min_total,omitempty"`
	MaxTotal           float64   `url:"max_total,omitempty"`
	CustomerID         int       `url:"customer_id,omitempty"`
	Email              string    `url:"email,omitempty"`
	StatusID           int       `url:"status_id,omitempty"`
	CartID             string    `url:"cart_id,omitempty"`
	PaymentMethod      string    `url:"payment_method,omitempty"`
	MinDateCreated     time.Time `url:"-"`
	MaxDateCreated     time.Time `url:"-"`
	MinDateModified    time.Time `url:"-"`
	MaxDateModified    time.Time `url:"-"`
	Page               int       `url:"page,omitempty"`
	Limit              int       `url:"limit,omitempty"`
	Sort               string    `url:"sort,omitempty"`
	IsDeleted          bool      `url:"is_deleted,omitempty"`
	MinDateCreatedRaw  string    `url:"min_date_created,omitempty"`
	MaxDateCreatedRaw  string    `url:"max_date_created,omitempty"`
	MinDateModifiedRaw string    `url:"min_date_modified,omitempty"`
	MaxDateModifiedRaw string    `url:"max_date_modified,omitempty"`
}

// GetRawQuery gets the struct in query string form
func (q OrderQuery) GetRawQuery() (raw string, err error) {
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
	raw = v.Encode()
	return
}

// EagerGet - attempts to unmarshal a url into an interface, preferably one intended to unmarshal the json body of that url.
func (r Resource) EagerGet(s *BCClient, i interface{}) error {
	url := r.URL
	body, err := s.GetBody(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, i)
	return err
}

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
	url := fmt.Sprintf(baseURL+"%s/v2/orders/count", s.StoreKey)

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
			url := fmt.Sprintf(baseURL+"%s/v2/orders/%d/shipments", s.StoreKey, o.ID)
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
	url := fmt.Sprintf(baseURL+"%s/v2/orders/", s.StoreKey)

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
	url := fmt.Sprintf(baseURL+"%s/v2/orders/%s", s.StoreKey, orderID)

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
