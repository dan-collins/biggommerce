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

type Status []StatusElement

type StatusElement struct {
	CustomLabel       *string `json:"custom_label,omitempty"`
	ID                *int64  `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Order             *int64  `json:"order,omitempty"`
	SystemDescription *string `json:"system_description,omitempty"`
	SystemLabel       *string `json:"system_label,omitempty"`
}

// OrderCount is a struct that represents the return body of BigCommerce GET /orders/count
//
// StatusCounts should be sorted based on the sort order in bigcommerce status
type OrderCount struct {
	StatusCounts []StatusCount `json:"statuses"`
	Count        int           `json:"count"`
}
