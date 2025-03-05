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
	completion, err := llms.GenerateFromSinglePrompt(
		context.Background(), llm, "hello deepseek",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(completion)
}
