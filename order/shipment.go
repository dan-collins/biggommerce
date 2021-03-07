package order

import "github.com/dan-collins/biggommerce/primative"

// Shipment is a struct that represents BC shipment
type Shipment struct {
	ID                   int64               `json:"id"`
	OrderID              int64               `json:"order_id"`
	CustomerID           int64               `json:"customer_id"`
	OrderAddressID       int64               `json:"order_address_id"`
	DateCreated          primative.BCDate    `json:"date_created"`
	TrackingNumber       string              `json:"tracking_number"`
	MerchantShippingCost float64             `json:"merchant_shipping_cost,string"`
	ShippingMethod       string              `json:"shipping_method"`
	Comments             string              `json:"comments"`
	ShippingProvider     string              `json:"shipping_provider"`
	TrackingCarrier      string              `json:"tracking_carrier"`
	TrackingLink         string              `json:"tracking_link"`
	BillingAddress       Address             `json:"billing_address"`
	ShippingAddress      Address             `json:"shipping_address"`
	Items                []ShipmentOrderItem `json:"items"`
}

// ShipmentOrderItem is a struct that represents BC orderitem
type ShipmentOrderItem struct {
	OrderProductID int64 `json:"order_product_id"`
	ProductID      int64 `json:"product_id"`
	Quantity       int64 `json:"quantity"`
}
