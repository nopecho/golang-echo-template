package http

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"net/http"
	"time"
)

type ANY map[string]any

type Client struct {
	host   string
	option *ClientOption
}

type ClientOption struct {
	client *http.Client
}

func DefaultOption() *ClientOption {
	transport := &http.Transport{
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		DisableKeepAlives:   false,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}
	return &ClientOption{
		client: &http.Client{
			Transport: transport,
		},
	}
}

func NewClient(host string, option *ClientOption) *Client {
	if option == nil {
		option = DefaultOption()
	}
	return &Client{
		host:   host,
		option: option,
	}
}

func (c *Client) GetJson(path string) (map[string]interface{}, error) {
	endpoint := c.makeEndpoint(path)
	resp, err := c.option.client.Get(endpoint)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	response, err := c.parseResponse(resp)
	return response, err
}

func (c *Client) GetRaw(path string) ([]byte, error) {
	endpoint := c.makeEndpoint(path)
	resp, err := c.option.client.Get(endpoint)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return res, err
}

func (c *Client) PostJson(path string, data interface{}) (map[string]interface{}, error) {
	endpoint := c.makeEndpoint(path)
	req, err := json.Marshal(data)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	resp, err := c.option.client.Post(endpoint, "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Err(err)
		return nil, err
	}

	response, err := c.parseResponse(resp)
	return response, err
}

func (c *Client) PatchJson(path string, data interface{}) (map[string]interface{}, error) {
	endpoint := c.makeEndpoint(path)
	req, err := json.Marshal(data)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	request, err := c.newJsonRequest("PATCH", endpoint, req)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	resp, err := c.option.client.Do(request)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	response, err := c.parseResponse(resp)
	return response, err
}

func (c *Client) newJsonRequest(method, endpoint string, req []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

func (c *Client) makeEndpoint(path string) string {
	return c.host + path
}

func (c *Client) parseResponse(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return nil, nil
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	var response map[string]interface{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return response, nil
}
