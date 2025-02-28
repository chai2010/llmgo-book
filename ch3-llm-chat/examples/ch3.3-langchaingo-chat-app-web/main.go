package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	s := NewLLMChatServer()
	s.Run("localhost:8080")
}

//go:embed static
var embedStaticFS embed.FS

type LLMChatServer struct {
	fs fs.FS
}

func NewLLMChatServer() *LLMChatServer {
	fs, err := fs.Sub(embedStaticFS, "static")
	if err != nil {
		panic(err)
	}

	p := &LLMChatServer{fs: fs}
	return p
}

func (p *LLMChatServer) Run(addr string) error {
	fmt.Println("listen on http://" + addr)
	startTime := time.Now()
	return http.ListenAndServe(addr,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method, r.URL.Path)

			switch {
			case r.URL.Path == "/":
				p.indexHandler(w, r)
			case r.URL.Path == "/run":
				p.runHandler(w, r)
			case strings.HasPrefix(r.URL.Path, "/static/"):
				relpath := strings.TrimPrefix(r.URL.Path, "/static/")
				data, err := fs.ReadFile(p.fs, relpath)
				if err != nil {
					http.NotFound(w, r)
					return
				}

				http.ServeContent(w, r, r.URL.Path, startTime, bytes.NewReader(data))

			default:
				http.NotFound(w, r)
			}
		}),
	)
}

func (p *LLMChatServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := fs.ReadFile(p.fs, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (p *LLMChatServer) runHandler(w http.ResponseWriter, r *http.Request) {
	prompt := struct {
		Input string `json:"input"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := LLMChat(prompt.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("input:", prompt.Input)
	fmt.Println("response:", resp)

	json.NewEncoder(w).Encode(map[string]string{
		"input":    prompt.Input,
		"response": resp,
	})
}

func LLMChat(prompt string) (string, error) {
	llm, err := ollama.New(ollama.WithModel("deepseek-r1:1.5b"))
	if err != nil {
		return "", err
	}

	return llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
}
