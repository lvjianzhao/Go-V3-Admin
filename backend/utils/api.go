package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func MakeAPIRequest(isAuth bool, url, method, body, basic string) ([]byte, error) {
	// 创建自定义的Transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建 HTTP 请求客户端
	client := &http.Client{
		Transport: transport,
	}

	// 创建请求体
	reqBody := bytes.NewBuffer([]byte(body))

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	if isAuth {
		// 添加基本认证头部
		req.Header.Set("Authorization", basic)
	}

	// 设置content-type
	req.Header.Set("Content-Type", "application/json")

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d, response body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
