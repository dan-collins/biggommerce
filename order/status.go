package order

//StatusCount is a struct that represents the individual statuses and count returned by BigCommerce GET /orders/count
type StatusCount struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SystemLabel string `json:"system_label"`
	CustomLabel string `json:"custom_label"`
	SystemDesc  string `json:"system_description"`
	Count       int    `json:"count"`
	SortOrder   int    `json:"sort_order"`
}

// OrderCount is a struct that represents the return body of BigCommerce GET /orders/count
type OrderCount struct {
	StatusCounts []StatusCount `json:"statuses"`
	Count        int           `json:"count"`
}
