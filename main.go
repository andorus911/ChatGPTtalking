package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
  // do not forget to create .env file!
	tokenKey := os.Getenv("OPENAI_API_KEY")

  client := openai.NewClient(tokenKey)
  resp, err := client.CreateChatCompletion(
    context.Background(),
    openai.ChatCompletionRequest{
      Model: openai.GPT3Dot5Turbo,
      Messages: []openai.ChatCompletionMessage{
        {
          Role:    openai.ChatMessageRoleUser,
          Content: "Hello!",
        },
      },
    },
  )

  if err != nil {
    fmt.Printf("ChatCompletion error: %v\n", err)
    return
  }

  fmt.Println(resp.Choices[0].Message.Content)
}
