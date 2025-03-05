package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		panic(err)
	}

	requestContent := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "你好，今天怎么样？"),
		llms.TextParts(llms.ChatMessageTypeSystem, "你好呀！我今天很好，谢谢！"),
		llms.TextParts(llms.ChatMessageTypeHuman, "你做了什么？"),
		llms.TextParts(llms.ChatMessageTypeSystem, "我今天一直在和你聊天！"),
		llms.TextParts(llms.ChatMessageTypeHuman, "那我们继续聊吧！"),
	}

	completion, err := llm.GenerateContent(context.Background(), requestContent)
	if err != nil {
		panic(err)
	}
	if len(completion.Choices) == 0 {
		panic("no response")
	}

	fmt.Println(completion.Choices[0].Content)
}
