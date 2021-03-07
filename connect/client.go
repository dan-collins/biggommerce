package connect

import (
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