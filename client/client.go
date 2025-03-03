package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a client for the API
type Client struct {
	BaseURL    string
	AuthToken  string
	HTTPClient *http.Client
}

// NewClient creates a new client with the given base URL and auth token
func NewClient(baseURL string, authToken string) *Client {
	return &Client{
		BaseURL:    baseURL,
		AuthToken:  authToken,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// SendRequest 发送通用 HTTP 请求
func (c *Client) SendRequest(method, endpoint string, body interface{}) (*http.Response, []byte, error) {
	url := c.BaseURL + endpoint
	var requestBody []byte
	var err error

	// 处理请求体
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
	}

	// 创建请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return resp, respBody, nil
}

// SendAuthRequest 发送带 Token 认证的 HTTP 请求
func (c *Client) SendAuthRequest(method, endpoint, token string, body interface{}) (*http.Response, []byte, error) {
	url := c.BaseURL + endpoint
	var requestBody []byte
	var err error

	// 处理请求体
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
	}

	// 创建请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return resp, respBody, nil
}
