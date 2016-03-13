package instagram

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Endpoint
	Endpoint = "https://api.instagram.com/v1"
)

// Client The main Instagram API client interface
type Client struct {
	token  string
	count  int
	client *http.Client
}

// NewClient Initialize Instagram Client
func NewClient(token string, count int) (*Client, error) {
	if token == "" {
		return nil, errors.New("Missing access token")
	}
	client := Client{}
	client.token = token
	client.count = count
	client.client = new(http.Client)
	return &client, nil
}

// PerformRequest is a Helper function to perform api call to Instagram server
// and provide common exception checking
func (c *Client) PerformRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, errors.New("Missing body")
	}

	return body, err
}

// GetRecentMedia show the recent media of the Instagram account based on the access token
func (c *Client) GetRecentMedia() ([]byte, error) {
	url := "/users/self/media/recent/"
	url += fmt.Sprintf("?count=%v&access_token=%s", c.count, c.token)
	req, err := http.NewRequest("GET", Endpoint+url, nil)
	// fmt.Println(url)
	if err != nil {
		return nil, err
	}
	resp, err := c.PerformRequest(req)
	// fmt.Println(string(resp))
	return resp, err
}
