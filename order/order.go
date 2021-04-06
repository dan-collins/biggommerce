package order

import (
	"time"

	"github.com/dan-collins/biggommerce/primative"
)

// Order is a struct that represents the return body of BigCommerce GET /orders
type Order struct {
	ID                                      int64              `json:"id,omitempty"`
	CustomerID                              int64              `json:"customer_id,omitempty"`
	DateCreated                             primative.BCDate   `json:"date_created,omitempty"`
	DateModified                            primative.BCDate   `json:"date_modified,omitempty"`
	DateShipped                             primative.BCDate   `json:"date_shipped,omitempty"`
	StatusID                                int64              `json:"status_id,omitempty"`
	Status                                  string             `json:"status,omitempty"`
	SubtotalExTax                           float64            `json:"subtotal_ex_tax,string"`
	SubtotalIncTax                          float64            `json:"subtotal_inc_tax,string"`
	SubtotalTax                             float64            `json:"subtotal_tax,string"`
	BaseShippingCost                        float64            `json:"base_shipping_cost,string"`
	ShippingCostExTax                       float64            `json:"shipping_cost_ex_tax,string"`
	ShippingCostIncTax                      float64            `json:"shipping_cost_inc_tax,string"`
	ShippingCostTax                         float64            `json:"shipping_cost_tax,string"`
	ShippingCostTaxClassID                  int64              `json:"shipping_cost_tax_class_id,omitempty"`
	BaseHandlingCost                        float64            `json:"base_handling_cost,string"`
	HandlingCostExTax                       float64            `json:"handling_cost_ex_tax,string"`
	HandlingCostIncTax                      float64            `json:"handling_cost_inc_tax,string"`
	HandlingCostTax                         float64            `json:"handling_cost_tax,string"`
	HandlingCostTaxClassID                  int64              `json:"handling_cost_tax_class_id,omitempty"`
	BaseWrappingCost                        float64            `json:"base_wrapping_cost,string"`
	WrappingCostExTax                       float64            `json:"wrapping_cost_ex_tax,string"`
	WrappingCostIncTax                      float64            `json:"wrapping_cost_inc_tax,string"`
	WrappingCostTax                         float64            `json:"wrapping_cost_tax,string"`
	WrappingCostTaxClassID                  int64              `json:"wrapping_cost_tax_class_id,omitempty"`
	TotalExTax                              float64            `json:"total_ex_tax,string"`
	TotalIncTax                             float64            `json:"total_inc_tax,string"`
	TotalTax                                float64            `json:"total_tax,string"`
	ItemsTotal                              int64              `json:"items_total,omitempty"`
	ItemsShipped                            int64              `json:"items_shipped,omitempty"`
	PaymentMethod                           string             `json:"payment_method,omitempty"`
	PaymentProviderID                       string             `json:"payment_provider_id,omitempty"`
	PaymentStatus                           string             `json:"payment_status,omitempty"`
	RefundedAmount                          float64            `json:"refunded_amount,string"`
	OrderIsDigital                          bool               `json:"order_is_digital,omitempty"`
	StoreCreditAmount                       float64            `json:"store_credit_amount,string"`
	GiftCertificateAmount                   float64            `json:"gift_certificate_amount,string"`
	IPAddress                               string             `json:"ip_address,omitempty"`
	GeoipCountry                            string             `json:"geoip_country,omitempty"`
	GeoipCountryIso2                        string             `json:"geoip_country_iso2,omitempty"`
	CurrencyID                              int64              `json:"currency_id,omitempty"`
	CurrencyCode                            string             `json:"currency_code,omitempty"`
	CurrencyExchangeRate                    string             `json:"currency_exchange_rate,omitempty"`
	DefaultCurrencyID                       int64              `json:"default_currency_id,omitempty"`
	DefaultCurrencyCode                     string             `json:"default_currency_code,omitempty"`
	StaffNotes                              string             `json:"staff_notes,omitempty"`
	CustomerMessage                         string             `json:"customer_message,omitempty"`
	DiscountAmount                          float64            `json:"discount_amount,string"`
	CouponDiscount                          float64            `json:"coupon_discount,string"`
	ShippingAddressCount                    int64              `json:"shipping_address_count,omitempty"`
	IsDeleted                               bool               `json:"is_deleted,omitempty"`
	EbayOrderID                             string             `json:"ebay_order_id,omitempty"`
	CartID                                  string             `json:"cart_id,omitempty"`
	BillingAddress                          Address            `json:"billing_address,omitempty"`
	IsEmailOptIn                            bool               `json:"is_email_opt_in,omitempty"`
	CreditCardType                          interface{}        `json:"credit_card_type"`
	OrderSource                             string             `json:"order_source,omitempty"`
	ChannelID                               int64              `json:"channel_id,omitempty"`
	ExternalSource                          interface{}        `json:"external_source"`
	ProductResource                         primative.Resource `json:"products,omitempty"`
	Products                                []OrderProduct
	ShippingResource                        primative.Resource `json:"shipping_addresses,omitempty"`
	ShippingAddresses                       []ShippingAddress
	CouponResource                          primative.Resource `json:"coupons,omitempty"`
	Coupons                                 []Coupon
	Shipments                               []Shipment
	ExternalID                              interface{} `json:"external_id"`
	ExternalMerchantID                      interface{} `json:"external_merchant_id"`
	TaxProviderID                           string      `json:"tax_provider_id,omitempty"`
	StoreDefaultCurrencyCode                string      `json:"store_default_currency_code,omitempty"`
	StoreDefaultToTransactionalExchangeRate string      `json:"store_default_to_transactional_exchange_rate,omitempty"`
	CustomStatus                            string      `json:"custom_status,omitempty"`
}

// Query struct to handle orders endpoint search query params, if you want orders with a status of 0 ("incomplete" in BC)
// you should set StatusIDIsZero to true, otherwise it will be ignored as a zero value when building the REST query
type Query struct {
	MinID              int       `url:"min_id,omitempty"`
	MaxID              int       `url:"max_id,omitempty"`
	MinTotal           float64   `url:"min_total,omitempty"`
	MaxTotal           float64   `url:"max_total,omitempty"`
	CustomerID         int       `url:"customer_id,omitempty"`
	Email              string    `url:"email,omitempty"`
	StatusID           int       `url:"status_id,omitempty"`
	StatusIDIsZero     bool      `url:"-"`
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
