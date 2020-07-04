package webdriver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host string
	Port string
}

func NewClient(host string, port string) *Client {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "4444"
	}

	return &Client{
		Host: host,
		Port: port,
	}
}

func (c *Client) URL() string {
	return fmt.Sprintf("http://%s:%s", c.Host, c.Port)
}

func (c *Client) Endpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.URL(), endpoint)
}

func (c *Client) Post(endpoint string, payload []byte) ([]byte, error) {
	resp, err := http.Post(c.Endpoint(endpoint), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	resp, err := http.Get(c.Endpoint(endpoint))
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (c *Client) Delete(endpoint string) ([]byte, error) {
	resp, err := http.NewRequest(http.MethodDelete, c.Endpoint(endpoint), nil)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
