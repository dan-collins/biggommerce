package primative

import "github.com/dan-collins/biggommerce/connect"

// Resource is general struct for resource url and type found in many returned objects from bigcommerce
type Resource struct {
	URL      string
	Resource string
}

// EagerGet - attempts to unmarshal a resource url into an interface, preferably one intended to unmarshal the json body of that url.
func (r Resource) EagerGet(s connect.Client, i interface{}) error {
	url := r.URL
	return s.GetAndUnmarshalRaw(url, &i)
}
