package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a company branding design wizard."),
		llms.TextParts(llms.ChatMessageTypeHuman, "What would be a good company name a company that makes colorful socks?"),
	}

	callOptions := []llms.CallOption{
		llms.WithMaxLength(50),
		//llms.WithMinLength(10),
		//llms.WithMaxTokens(50),
	}

	resp, err := llm.GenerateContent(ctx, content, callOptions...)

	if err != nil {
		log.Fatal(err)
	}

	//lenth of the response
	fmt.Println(len(resp.Choices[0].Content))
	fmt.Println(resp.Choices[0].Content)
}
