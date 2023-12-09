package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type ReqBody struct {
	Msg string `json:"msg"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	CheckErrorFatal("Error loading .env file", err)

	// set openAI constants
	openAIConstants := OpenAIConstants{
		Key:   os.Getenv("OPENAI_KEY"),
		Model: openai.GPT3Dot5Turbo,
	}

	// msg := "what is the tallest mountain in the world?"
	InitWaku()

	r := InitServer(openAIConstants)
	StartServer(r)
}
