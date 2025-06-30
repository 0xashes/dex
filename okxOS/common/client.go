package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/joho/godotenv"
)

// Client 通用客户端
type Client struct {
	BaseURL    string
	APIKey     string
	SecretKey  string
	Passphrase string
	HTTPClient *http.Client
}

// NewClient 初始化客户端
func NewClient() *Client {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Failed to load .env file: %v. Using default environment variables.", err)
	}

	// 获取代理设置（例如 http://127.0.0.1:7890）
	proxyURLStr := os.Getenv("HTTP_PROXY")
	var transport *http.Transport

	if proxyURLStr != "" {
		proxyURL, err := url.Parse(proxyURLStr)
		if err != nil {
			log.Fatalf("Invalid proxy URL: %v", err)
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = http.DefaultTransport.(*http.Transport).Clone()
	}

	// 初始化 HTTP 客户端
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	// 读取环境变量
	cfg := &Client{
		BaseURL:    "https://web3.okx.com",
		APIKey:     os.Getenv("OKX_OS_API_KEY"),
		SecretKey:  os.Getenv("OKX_OS_SECRET_KEY"),
		Passphrase: os.Getenv("OKX_OS_PASSPHRASE"),
		HTTPClient: httpClient,
	}

	// 验证必填字段
	if cfg.BaseURL == "" || cfg.APIKey == "" || cfg.SecretKey == "" || cfg.Passphrase == "" {
		log.Fatal("Error: Missing required environment variables (BASE_URL, API_KEY, SECRET_KEY, PASSPHRASE)")
	}

	return cfg
}

// DoRequest 执行 HTTP 请求
func (c *Client) DoRequest(method, endpoint string, params interface{}, response interface{}) error {
	// endpoint = strings.TrimPrefix(endpoint, "/")
	query := ""
	bodyStr := ""

	// 处理查询参数（GET 请求）
	if method == "GET" && params != nil {
		queryValues, err := paramsToQuery(params)
		if err != nil {
			log.Printf("Failed to encode query params: %v", err)
			return err
		}
		query = queryValues.Encode()
		if query != "" {
			query = "?" + query
		}
	} else if params != nil {
		// POST 等请求，参数作为 JSON body
		body, err := json.Marshal(params)
		if err != nil {
			log.Printf("Failed to marshal params: %v", err)
			return err
		}
		bodyStr = string(body)
	}

	fullURL := c.BaseURL + endpoint + query
	requestPath := endpoint + query

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 15 * time.Second
	bo.MaxInterval = 5 * time.Second

	var lastErr error
	err := backoff.Retry(func() error {
		var body []byte
		if bodyStr != "" {
			body = []byte(bodyStr)
		}

		req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return backoff.Permanent(err)
		}

		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		sign := c.generateSignature(timestamp, method, requestPath, bodyStr)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("OK-ACCESS-KEY", c.APIKey)
		req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
		req.Header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
		req.Header.Set("OK-ACCESS-SIGN", sign)

		resp, err := c.HTTPClient.Do(req)

		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			log.Printf("Decode response failed: %v", err)
			return backoff.Permanent(err)
		}

		return nil
	}, bo)

	if err != nil {
		log.Printf("Request failed after retries: %v", err)
		lastErr = err
	}
	return lastErr
}

// RequestRaw 用于调试：执行请求并返回原始响应体（方便构造结构体）
func (c *Client) RequestRaw(method, endpoint string, params interface{}) ([]byte, error) {
	query := ""
	bodyStr := ""

	if method == "GET" && params != nil {
		queryValues, err := paramsToQuery(params)
		if err != nil {
			log.Printf("Failed to encode query params: %v", err)
			return nil, err
		}
		query = queryValues.Encode()
		if query != "" {
			query = "?" + query
		}
	} else if params != nil {
		body, err := json.Marshal(params)
		if err != nil {
			log.Printf("Failed to marshal params: %v", err)
			return nil, err
		}
		bodyStr = string(body)
	}

	fullURL := c.BaseURL + endpoint + query
	requestPath := endpoint + query

	var body []byte
	if bodyStr != "" {
		body = []byte(bodyStr)
	}
	// fmt.Printf("bodyStr: %v\n", bodyStr)

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	sign := c.generateSignature(timestamp, method, requestPath, bodyStr)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OK-ACCESS-KEY", c.APIKey)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
	req.Header.Set("OK-ACCESS-SIGN", sign)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// paramsToQuery 将参数转换为查询字符串
func paramsToQuery(params interface{}) (url.Values, error) {
	values := url.Values{}
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	for k, v := range m {
		if v != nil && v != "" {
			values.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return values, nil
}

// generateSignature 生成 OKX API 签名
func (c *Client) generateSignature(timestamp, method, requestPath, body string) string {
	message := timestamp + strings.ToUpper(method) + requestPath + body
	h := hmac.New(sha256.New, []byte(c.SecretKey))
	h.Write([]byte(message))
	signature := h.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(signature)
	return encoded
}
