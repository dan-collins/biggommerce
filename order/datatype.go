package order

import "github.com/dan-collins/biggommerce/primative"

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
	Amount float64     `json:"amount,string,omitempty"`
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
	ID                   int64            `json:"id"`
	OrderID              int64            `json:"order_id"`
	CustomerID           int64            `json:"customer_id"`
	OrderAddressID       int64            `json:"order_address_id"`
	DateCreated          primative.BCDate `json:"date_created"`
	TrackingNumber       string           `json:"tracking_number"`
	MerchantShippingCost float64          `json:"merchant_shipping_cost,string"`
	ShippingMethod       string           `json:"shipping_method"`
	Comments             string           `json:"comments"`
	ShippingProvider     string           `json:"shipping_provider"`
	TrackingCarrier      string           `json:"tracking_carrier"`
	TrackingLink         string           `json:"tracking_link"`
	BillingAddress       Address          `json:"billing_address"`
	ShippingAddress      Address          `json:"shipping_address"`
	Items                []OrderItem      `json:"items"`
}

// OrderItem struct representing BC orderitem
type OrderItem struct {
	OrderProductID int64 `json:"order_product_id"`
	ProductID      int64 `json:"product_id"`
	Quantity       int64 `json:"quantity"`
}

// Order struct representing the return body of BigCommerce GET /orders
type Order struct {
	ID                                      int64            `json:"id,omitempty"`
	CustomerID                              int64            `json:"customer_id,omitempty"`
	DateCreated                             primative.BCDate `json:"date_created,omitempty"`
	DateModified                            primative.BCDate `json:"date_modified,omitempty"`
	DateShipped                             primative.BCDate `json:"date_shipped,omitempty"`
	StatusID                                int64            `json:"status_id,omitempty"`
	Status                                  string           `json:"status,omitempty"`
	SubtotalExTax                           float64          `json:"subtotal_ex_tax,string"`
	SubtotalIncTax                          float64          `json:"subtotal_inc_tax,string"`
	SubtotalTax                             float64          `json:"subtotal_tax,string"`
	BaseShippingCost                        float64          `json:"base_shipping_cost,string"`
	ShippingCostExTax                       float64          `json:"shipping_cost_ex_tax,string"`
	ShippingCostIncTax                      float64          `json:"shipping_cost_inc_tax,string"`
	ShippingCostTax                         float64          `json:"shipping_cost_tax,string"`
	ShippingCostTaxClassID                  int64            `json:"shipping_cost_tax_class_id,omitempty"`
	BaseHandlingCost                        float64          `json:"base_handling_cost,string"`
	HandlingCostExTax                       float64          `json:"handling_cost_ex_tax,string"`
	HandlingCostIncTax                      float64          `json:"handling_cost_inc_tax,string"`
	HandlingCostTax                         float64          `json:"handling_cost_tax,string"`
	HandlingCostTaxClassID                  int64            `json:"handling_cost_tax_class_id,omitempty"`
	BaseWrappingCost                        float64          `json:"base_wrapping_cost,string"`
	WrappingCostExTax                       float64          `json:"wrapping_cost_ex_tax,string"`
	WrappingCostIncTax                      float64          `json:"wrapping_cost_inc_tax,string"`
	WrappingCostTax                         float64          `json:"wrapping_cost_tax,string"`
	WrappingCostTaxClassID                  int64            `json:"wrapping_cost_tax_class_id,omitempty"`
	TotalExTax                              float64          `json:"total_ex_tax,string"`
	TotalIncTax                             float64          `json:"total_inc_tax,string"`
	TotalTax                                float64          `json:"total_tax,string"`
	ItemsTotal                              int64            `json:"items_total,omitempty"`
	ItemsShipped                            int64            `json:"items_shipped,omitempty"`
	PaymentMethod                           string           `json:"payment_method,omitempty"`
	PaymentProviderID                       string           `json:"payment_provider_id,omitempty"`
	PaymentStatus                           string           `json:"payment_status,omitempty"`
	RefundedAmount                          float64          `json:"refunded_amount,string"`
	OrderIsDigital                          bool             `json:"order_is_digital,omitempty"`
	StoreCreditAmount                       float64          `json:"store_credit_amount,string"`
	GiftCertificateAmount                   float64          `json:"gift_certificate_amount,string"`
	IPAddress                               string           `json:"ip_address,omitempty"`
	GeoipCountry                            string           `json:"geoip_country,omitempty"`
	GeoipCountryIso2                        string           `json:"geoip_country_iso2,omitempty"`
	CurrencyID                              int64            `json:"currency_id,omitempty"`
	CurrencyCode                            string           `json:"currency_code,omitempty"`
	CurrencyExchangeRate                    string           `json:"currency_exchange_rate,omitempty"`
	DefaultCurrencyID                       int64            `json:"default_currency_id,omitempty"`
	DefaultCurrencyCode                     string           `json:"default_currency_code,omitempty"`
	StaffNotes                              string           `json:"staff_notes,omitempty"`
	CustomerMessage                         string           `json:"customer_message,omitempty"`
	DiscountAmount                          float64          `json:"discount_amount,string"`
	CouponDiscount                          float64          `json:"coupon_discount,string"`
	ShippingAddressCount                    int64            `json:"shipping_address_count,omitempty"`
	IsDeleted                               bool             `json:"is_deleted,omitempty"`
	EbayOrderID                             string           `json:"ebay_order_id,omitempty"`
	CartID                                  string           `json:"cart_id,omitempty"`
	BillingAddress                          Address          `json:"billing_address,omitempty"`
	IsEmailOptIn                            bool             `json:"is_email_opt_in,omitempty"`
	CreditCardType                          interface{}      `json:"credit_card_type"`
	OrderSource                             string           `json:"order_source,omitempty"`
	ChannelID                               int64            `json:"channel_id,omitempty"`
	ExternalSource                          interface{}      `json:"external_source"`
	ProductResource                         Resource         `json:"products,omitempty"`
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
	BasePrice            float64           `json:"base_price,string"`
	PriceExTax           float64           `json:"price_ex_tax,string"`
	PriceIncTax          float64           `json:"price_inc_tax,string"`
	PriceTax             float64           `json:"price_tax,string"`
	BaseTotal            float64           `json:"base_total,string"`
	TotalExTax           float64           `json:"total_ex_tax,string"`
	TotalIncTax          float64           `json:"total_inc_tax,string"`
	TotalTax             float64           `json:"total_tax,string"`
	Weight               float64           `json:"weight,string"`
	Quantity             int64             `json:"quantity,omitempty"`
	BaseCostPrice        float64           `json:"base_cost_price,string"`
	CostPriceIncTax      float64           `json:"cost_price_inc_tax,string"`
	CostPriceExTax       float64           `json:"cost_price_ex_tax,string"`
	CostPriceTax         float64           `json:"cost_price_tax,string"`
	IsRefunded           bool              `json:"is_refunded,omitempty"`
	QuantityRefunded     int64             `json:"quantity_refunded,omitempty"`
	RefundAmount         float64           `json:"refund_amount,string"`
	ReturnID             int64             `json:"return_id,omitempty"`
	WrappingName         string            `json:"wrapping_name,omitempty"`
	BaseWrappingCost     float64           `json:"base_wrapping_cost,string"`
	WrappingCostExTax    float64           `json:"wrapping_cost_ex_tax,string"`
	WrappingCostIncTax   float64           `json:"wrapping_cost_inc_tax,string"`
	WrappingCostTax      float64           `json:"wrapping_cost_tax,string"`
	WrappingMessage      string            `json:"wrapping_message,omitempty"`
	QuantityShipped      int64             `json:"quantity_shipped,omitempty"`
	FixedShippingCost    float64           `json:"fixed_shipping_cost,string"`
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
