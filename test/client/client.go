package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	URL    string
	client *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		URL:    url,
		client: &http.Client{},
	}
}

func (c *Client) Get(path string, params map[string]string, responseData interface{}) (int, error) {
	url := c.buildURL(path, params)
	return c.Do(http.MethodGet, url, nil, responseData)
}

func (c *Client) Post(path string, requestData interface{}, responseData interface{}) (int, error) {
	url := c.URL + path
	return c.Do(http.MethodPost, url, requestData, responseData)
}

func (c *Client) buildURL(path string, params map[string]string) string {
	u, _ := url.Parse(c.URL + path)
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *Client) Do(method, url string, input interface{}, result interface{}) (int, error) {

	var requestBody []byte
	if input != nil {
		requestBody, _ = json.Marshal(input)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return 500, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if result != nil {
		err := json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return -1, err
		}
	}

	return resp.StatusCode, nil
}
