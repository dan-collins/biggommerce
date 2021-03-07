package order

// Coupon is a struct that represents the coupon detail objects that are part of the order
type Coupon struct {
	Amount   string  `json:"amount,omitempty"`
	Code     string  `json:"code,omitempty"`
	CouponID int64   `json:"coupon_id,omitempty"`
	Discount float64 `json:"discount,string"`
	ID       int64   `json:"id,omitempty"`
	OrderID  int64   `json:"order_id,omitempty"`
	Type     int64   `json:"type,omitempty"`
}

// CouponType will get you the text representation of the coupon type
func (c *Coupon) CouponType() string {
	types := map[int64]string{
		0: "per_item_discount",
		1: "percentage_discount",
		2: "per_total_discount",
		3: "shipping_discount",
		4: "free_shipping",
		5: "promotion",
	}
	if coupType, ok := types[c.Type]; ok {
		return coupType
	}
	return ""
}
