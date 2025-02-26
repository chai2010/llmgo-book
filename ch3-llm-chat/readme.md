# 3. LLM聊天机器人

LangChain是专门用于开发LLM驱动型应用程序的框架，而LangChain Go是LangChain框架的Go语言实现。现在我们尝试用LangChain Go连接Ollama运行的本地大模型，然后构建一个LLM聊天机器人。

## 3.1 连接Ollama

通过LangChainGo连接Ollama非常简单：

```go
package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, _ := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	completion, _ := llms.GenerateFromSinglePrompt(
		context.Background(), llm, "hello deepseek",
	)
	fmt.Println(completion)
}
```

本地执行结果如下：

```
$ go run .
<think>

</think>

Hello! How can I assist you today? 😊
```

## 3.2 命令行聊天程序

先构造一个命令行环境的聊天程序，每次对话不包含上下文信息。首先包装一个`LLMChat`函数：


```go
func LLMChat(prompt string) (string, error) {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		return "", err
	}

	return llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
}
```

然后再main函数中以交互的方式调用`LLMChat`函数：

```go
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
```

执行的效果如下：

```
$ go run .
请输入聊天对话内容，输入 '/bye' 退出。
>> hello deepseek
<think>

</think>

Hello! How can I assist you today? 😊
>> /bye
$
```

