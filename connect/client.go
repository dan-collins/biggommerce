package connect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseBCURL string = "https://api.bigcommerce.com/stores/"

//BCClient represents the client wrapper containing BC specific connection details
type BCClient struct {
	AuthToken  string
	AuthClient string
	StoreKey   string
	BaseURL    string
}

//NewClient create a new client wrapper based on BC connection details
func NewClient(authToken, authClient, storeKey string) *BCClient {
	return &BCClient{
		AuthToken:  authToken,
		AuthClient: authClient,
		StoreKey:   storeKey,
		BaseURL:    baseBCURL,
	}
}

// SetBaseURL will override the default (https://api.bigcommerce.com/stores/) base url of the client to the string passed in
func (s *BCClient) SetBaseURL(url string) {
	s.BaseURL = url
}

// DoRequest will call out to bigcommerce API for the passed in request, this is mostly used internal to the
// package but can be used to expand on the library
func (s *BCClient) DoRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-auth-token", s.AuthToken)
	req.Header.Add("x-auth-client", s.AuthClient)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// GetBody - gets the request body of the url
func (s *BCClient) GetBody(url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	body, err = s.DoRequest(req)
	return
}

// BuildUrl - gets the golang request out of the endpoint (e.g. /v2/orders/) needed for sending to the BC api
func (s *BCClient) BuildUrlRequest(endpoint string) (req *http.Request, err error) {
	url := fmt.Sprintf(s.BaseURL+"%s/%s", s.StoreKey, endpoint)

	req, err = http.NewRequest("GET", url, nil)
	return
}

// GetAndUnmarshal - gets the request body of a plain url and unmarshals to passed in struct pointer
//
// Example of the endpoint parameter would be "/v2/orders/" and the client will handle the store key and base url pieces
func (s *BCClient) GetAndUnmarshal(endpoint string, outData interface{}) error {
	req, err := s.BuildUrlRequest(endpoint)
	if err != nil {
		return err
	}

	return s.doUnmarshalling(req, outData)
}

func (s *BCClient) doUnmarshalling(req *http.Request, outData interface{}) error {
	res, err := s.DoRequest(req)
	if err != nil {
		return err
	}

	if len(res) > 0 {
		err = json.Unmarshal(res, outData)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAndUnmarshalRaw - gets the request body of a full plain bigcommerce url and unmarshals to passed in struct pointer
//
// Example of the fullEndpoint parameter would be "https://api.bigcommerce.com/stores/{{YOUR-STORE-KEY}}/v2/orders/12039/products"
// the client will not manipulate the endpoint in any way.
// This function is used by things like Resource.EagerGet
func (s *BCClient) GetAndUnmarshalRaw(fullEndpoint string, outData interface{}) error {
	req, err := http.NewRequest("GET", fullEndpoint, nil)
	if err != nil {
		return err
	}
	return s.doUnmarshalling(req, outData)
}

// GetAndUnmarshalWithQuery - gets the request body of the url with a query string added on and unmarshals to passed in struct pointer
//
// Example of the endpoint parameter would be "/v2/orders/count/" and the client will handle the store key and base url pieces
func (s *BCClient) GetAndUnmarshalWithQuery(endpoint string, rawQuery string, outData interface{}) error {
	req, err := s.BuildUrlRequest(endpoint)
	if err != nil {
		return err
	}
	req.URL.RawQuery = rawQuery

	return s.doUnmarshalling(req, outData)
}

type Client interface {
	SetBaseURL(url string)
	DoRequest(req *http.Request) ([]byte, error)
	GetBody(url string) (body []byte, err error)
	BuildUrlRequest(endpoint string) (req *http.Request, err error)
	GetAndUnmarshal(endpoint string, outData interface{}) error
	GetAndUnmarshalRaw(fullEndpoint string, outData interface{}) error
	GetAndUnmarshalWithQuery(endpoint string, rawQuery string, outData interface{}) error
}
