package order

// Address is a struct that represents a BigCommerce address resource
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

// ShippingAddress is a struct that represents a BigCommerce shipping address resource
type ShippingAddress struct {
	Address
	ShippingMethod string `json:"shipping_method"`
}

// FormField struct that represents a BigCommerce address form name/value pair
type FormField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
