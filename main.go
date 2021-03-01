package biggommerce

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//BCClient represents the client wrapper containing BC specific connection details
type BCClient struct {
	AuthToken  string
	AuthClient string
	StoreKey   string
}

const baseURL string = "https://api.bigcommerce.com/stores/"

//NewClient create a new client wrapper based on BC connection details
func NewClient(authToken, authClient, storeKey string) *BCClient {
	return &BCClient{
		AuthToken:  authToken,
		AuthClient: authClient,
		StoreKey:   storeKey,
	}
}

func (s *BCClient) doRequest(req *http.Request) ([]byte, error) {
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
	if 200 != resp.StatusCode {
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
	body, err = s.doRequest(req)
	return
}
