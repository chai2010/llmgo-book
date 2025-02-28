package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

//go:embed llama.png
var llamaImageData []byte

func main() {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		panic(err)
	}

	requestContent := []llms.MessageContent{
		llms.MessageContent{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.BinaryPart("image/png", llamaImageData),
			},
		},
		llms.MessageContent{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: "What's in this image?"},
			},
		},
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
