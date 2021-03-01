package model

import (
	"encoding/json"
	"time"

	"github.com/google/go-querystring/query"
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

type AppliedDiscount struct {
	ID     string      `json:"id,omitempty"`
	Amount string      `json:"amount,omitempty"`
	Name   string      `json:"name,omitempty"`
	Code   interface{} `json:"code"`
	Target string      `json:"target,omitempty"`
}

type ProductOption struct {
	ID              int64  `json:"id,omitempty"`
	OptionID        int64  `json:"option_id,omitempty"`
	OrderProductID  int64  `json:"order_product_id,omitempty"`
	ProductOptionID int64  `json:"product_option_id,omitempty"`
	DisplayName     string `json:"display_name,omitempty"`
	DisplayValue    string `json:"display_value,omitempty"`
	Value           string `json:"value,omitempty"`
	Type            string `json:"type,omitempty"`
	Name            string `json:"name,omitempty"`
	DisplayStyle    string `json:"display_style,omitempty"`
}

// Address struct representing a BigCommerce address resource
type Address struct {
	FirstName   string      `json:"first_name,omitempty"`
	LastName    string      `json:"last_name,omitempty"`
	Company     string      `json:"company,omitempty"`
	Street1     string      `json:"street_1,omitempty"`
	Street2     string      `json:"street_2,omitempty"`
	City        string      `json:"city,omitempty"`
	State       string      `json:"state,omitempty"`
	Zip         string      `json:"zip,omitempty"`
	Country     string      `json:"country,omitempty"`
	CountryIso2 string      `json:"country_iso2,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Email       string      `json:"email,omitempty"`
	FormFields  []FormField `json:"form_fields,omitempty"`
}

// FormField struct that represents a BigCommerce address form name/value pair
type FormField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// ShippingAddress struct representing a BigCommerce shipping address resource
type ShippingAddress struct {
	Address
	ShippingMethod string `json:"shipping_method"`
}

// Shipment struct representing BC shipment
type Shipment struct {
	ID                   int64       `json:"id"`
	OrderID              int64       `json:"order_id"`
	CustomerID           int64       `json:"customer_id"`
	OrderAddressID       int64       `json:"order_address_id"`
	DateCreated          BCDate      `json:"date_created"`
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
	ID                                      int64       `json:"id,omitempty"`
	CustomerID                              int64       `json:"customer_id,omitempty"`
	DateCreated                             BCDate      `json:"date_created,omitempty"`
	DateModified                            BCDate      `json:"date_modified,omitempty"`
	DateShipped                             BCDate      `json:"date_shipped,omitempty"`
	StatusID                                int64       `json:"status_id,omitempty"`
	Status                                  string      `json:"status,omitempty"`
	SubtotalExTax                           string      `json:"subtotal_ex_tax,omitempty"`
	SubtotalIncTax                          string      `json:"subtotal_inc_tax,omitempty"`
	SubtotalTax                             string      `json:"subtotal_tax,omitempty"`
	BaseShippingCost                        string      `json:"base_shipping_cost,omitempty"`
	ShippingCostExTax                       string      `json:"shipping_cost_ex_tax,omitempty"`
	ShippingCostIncTax                      string      `json:"shipping_cost_inc_tax,omitempty"`
	ShippingCostTax                         string      `json:"shipping_cost_tax,omitempty"`
	ShippingCostTaxClassID                  int64       `json:"shipping_cost_tax_class_id,omitempty"`
	BaseHandlingCost                        string      `json:"base_handling_cost,omitempty"`
	HandlingCostExTax                       string      `json:"handling_cost_ex_tax,omitempty"`
	HandlingCostIncTax                      string      `json:"handling_cost_inc_tax,omitempty"`
	HandlingCostTax                         string      `json:"handling_cost_tax,omitempty"`
	HandlingCostTaxClassID                  int64       `json:"handling_cost_tax_class_id,omitempty"`
	BaseWrappingCost                        string      `json:"base_wrapping_cost,omitempty"`
	WrappingCostExTax                       string      `json:"wrapping_cost_ex_tax,omitempty"`
	WrappingCostIncTax                      string      `json:"wrapping_cost_inc_tax,omitempty"`
	WrappingCostTax                         string      `json:"wrapping_cost_tax,omitempty"`
	WrappingCostTaxClassID                  int64       `json:"wrapping_cost_tax_class_id,omitempty"`
	TotalExTax                              string      `json:"total_ex_tax,omitempty"`
	TotalIncTax                             string      `json:"total_inc_tax,omitempty"`
	TotalTax                                string      `json:"total_tax,omitempty"`
	ItemsTotal                              int64       `json:"items_total,omitempty"`
	ItemsShipped                            int64       `json:"items_shipped,omitempty"`
	PaymentMethod                           string      `json:"payment_method,omitempty"`
	PaymentProviderID                       string      `json:"payment_provider_id,omitempty"`
	PaymentStatus                           string      `json:"payment_status,omitempty"`
	RefundedAmount                          string      `json:"refunded_amount,omitempty"`
	OrderIsDigital                          bool        `json:"order_is_digital,omitempty"`
	StoreCreditAmount                       string      `json:"store_credit_amount,omitempty"`
	GiftCertificateAmount                   string      `json:"gift_certificate_amount,omitempty"`
	IPAddress                               string      `json:"ip_address,omitempty"`
	GeoipCountry                            string      `json:"geoip_country,omitempty"`
	GeoipCountryIso2                        string      `json:"geoip_country_iso2,omitempty"`
	CurrencyID                              int64       `json:"currency_id,omitempty"`
	CurrencyCode                            string      `json:"currency_code,omitempty"`
	CurrencyExchangeRate                    string      `json:"currency_exchange_rate,omitempty"`
	DefaultCurrencyID                       int64       `json:"default_currency_id,omitempty"`
	DefaultCurrencyCode                     string      `json:"default_currency_code,omitempty"`
	StaffNotes                              string      `json:"staff_notes,omitempty"`
	CustomerMessage                         string      `json:"customer_message,omitempty"`
	DiscountAmount                          string      `json:"discount_amount,omitempty"`
	CouponDiscount                          string      `json:"coupon_discount,omitempty"`
	ShippingAddressCount                    int64       `json:"shipping_address_count,omitempty"`
	IsDeleted                               bool        `json:"is_deleted,omitempty"`
	EbayOrderID                             string      `json:"ebay_order_id,omitempty"`
	CartID                                  string      `json:"cart_id,omitempty"`
	BillingAddress                          Address     `json:"billing_address,omitempty"`
	IsEmailOptIn                            bool        `json:"is_email_opt_in,omitempty"`
	CreditCardType                          interface{} `json:"credit_card_type"`
	OrderSource                             string      `json:"order_source,omitempty"`
	ChannelID                               int64       `json:"channel_id,omitempty"`
	ExternalSource                          interface{} `json:"external_source"`
	ProductResource                         Resource    `json:"products,omitempty"`
	Products                                []OrderProduct
	ShippingResource                        Resource `json:"shipping_addresses,omitempty"`
	ShippingAddresses                       []ShippingAddress
	Coupons                                 Resource    `json:"coupons,omitempty"`
	ExternalID                              interface{} `json:"external_id"`
	ExternalMerchantID                      interface{} `json:"external_merchant_id"`
	TaxProviderID                           string      `json:"tax_provider_id,omitempty"`
	StoreDefaultCurrencyCode                string      `json:"store_default_currency_code,omitempty"`
	StoreDefaultToTransactionalExchangeRate string      `json:"store_default_to_transactional_exchange_rate,omitempty"`
	CustomStatus                            string      `json:"custom_status,omitempty"`
}

// OrderProduct struct representing the product sub object that is part of the Order
type OrderProduct struct {
	ID                   int64             `json:"id,omitempty"`
	OrderID              int64             `json:"order_id,omitempty"`
	ProductID            int64             `json:"product_id,omitempty"`
	OrderAddressID       int64             `json:"order_address_id,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Sku                  string            `json:"sku,omitempty"`
	Upc                  string            `json:"upc,omitempty"`
	Type                 string            `json:"type,omitempty"`
	BasePrice            string            `json:"base_price,omitempty"`
	PriceExTax           string            `json:"price_ex_tax,omitempty"`
	PriceIncTax          string            `json:"price_inc_tax,omitempty"`
	PriceTax             string            `json:"price_tax,omitempty"`
	BaseTotal            string            `json:"base_total,omitempty"`
	TotalExTax           string            `json:"total_ex_tax,omitempty"`
	TotalIncTax          string            `json:"total_inc_tax,omitempty"`
	TotalTax             string            `json:"total_tax,omitempty"`
	Weight               string            `json:"weight,omitempty"`
	Quantity             int64             `json:"quantity,omitempty"`
	BaseCostPrice        string            `json:"base_cost_price,omitempty"`
	CostPriceIncTax      string            `json:"cost_price_inc_tax,omitempty"`
	CostPriceExTax       string            `json:"cost_price_ex_tax,omitempty"`
	CostPriceTax         string            `json:"cost_price_tax,omitempty"`
	IsRefunded           bool              `json:"is_refunded,omitempty"`
	QuantityRefunded     int64             `json:"quantity_refunded,omitempty"`
	RefundAmount         string            `json:"refund_amount,omitempty"`
	ReturnID             int64             `json:"return_id,omitempty"`
	WrappingName         string            `json:"wrapping_name,omitempty"`
	BaseWrappingCost     string            `json:"base_wrapping_cost,omitempty"`
	WrappingCostExTax    string            `json:"wrapping_cost_ex_tax,omitempty"`
	WrappingCostIncTax   string            `json:"wrapping_cost_inc_tax,omitempty"`
	WrappingCostTax      string            `json:"wrapping_cost_tax,omitempty"`
	WrappingMessage      string            `json:"wrapping_message,omitempty"`
	QuantityShipped      int64             `json:"quantity_shipped,omitempty"`
	FixedShippingCost    string            `json:"fixed_shipping_cost,omitempty"`
	EbayItemID           string            `json:"ebay_item_id,omitempty"`
	EbayTransactionID    string            `json:"ebay_transaction_id,omitempty"`
	OptionSetID          int64             `json:"option_set_id,omitempty"`
	ParentOrderProductID interface{}       `json:"parent_order_product_id,omitempty"`
	IsBundledProduct     bool              `json:"is_bundled_product,omitempty"`
	BinPickingNumber     string            `json:"bin_picking_number,omitempty"`
	ExternalID           interface{}       `json:"external_id,omitempty"`
	FulfillmentSource    string            `json:"fulfillment_source,omitempty"`
	AppliedDiscounts     []AppliedDiscount `json:"applied_discounts,omitempty"`
	ProductOptions       []ProductOption   `json:"product_options,omitempty"`
	ConfigurableFields   []interface{}     `json:"configurable_fields,omitempty"`
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
	if err != nil {
		return "", err
	}
	raw = v.Encode()
	return
}

type Client interface {
	GetBody(url string) ([]byte, error)
}

// EagerGet - attempts to unmarshal a url into an interface, preferably one intended to unmarshal the json body of that url.
func (r Resource) EagerGet(s Client, i interface{}) error {
	url := r.URL
	body, err := s.GetBody(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, i)
	return err
}
