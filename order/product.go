package order

// AppliedDiscount is a struct that represents the applied discounts included in the order product resource
type AppliedDiscount struct {
	ID     string      `json:"id,omitempty"`
	Amount float64     `json:"amount,string,omitempty"`
	Name   string      `json:"name,omitempty"`
	Code   interface{} `json:"code"`
	Target string      `json:"target,omitempty"`
}

// ProductOption is a struct that represents the product option detail objects that are part of the order product
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

// OrderProduct is a struct that represents the product sub object that is part of the Order
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
