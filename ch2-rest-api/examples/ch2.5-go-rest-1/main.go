package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 构造请求体
	requestData := map[string]any{
		"model": "deepseek-r1:1.5b",
		"messages": []map[string]any{
			{"role": "user", "content": "你好，今天怎么样？"},
			{"role": "assistant", "content": "你好呀！我今天很好，谢谢！"},
			{"role": "user", "content": "你做了什么？"},
			{"role": "assistant", "content": "我今天一直在和你聊天！"},
			{"role": "user", "content": "那我们继续聊吧！"},
		},
		"temperature": 0.7,
		"max_tokens":  150,
	}

	// 将请求体序列化为JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("无法序列化请求数据: %v", err)
	}

	// 创建 POST 请求
	url := "http://localhost:11434/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("无法创建请求: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	// 解析响应数据
	var result struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatalf("解码响应失败: %v", err)
	}

	fmt.Println(result.Choices[0].Message.Content)
}
