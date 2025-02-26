package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	// 提示用户
	fmt.Println("请输入聊天对话内容，输入 '/bye' 退出。")

	// 循环读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// 提示用户输入
		fmt.Print(">> ")

		// 读取一行输入
		scanner.Scan()
		input := scanner.Text()

		// 如果输入的是 'exit'，退出循环
		if strings.ToLower(input) == "/bye" {
			break
		}

		response, err := LLMChat(input)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		fmt.Println(response)
	}
}

func LLMChat(prompt string) (string, error) {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		return "", err
	}

	return llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
}
