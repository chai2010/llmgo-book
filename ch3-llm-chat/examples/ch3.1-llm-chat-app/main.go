package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 提示用户
	fmt.Println("请输入聊天对话内容，输入 '/bye' 退出。")

	// 循环读取输入
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

		// 输出用户输入
		fmt.Printf("你输入了: %s\n", input)
	}
}
