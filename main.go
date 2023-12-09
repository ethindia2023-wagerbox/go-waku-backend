package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Eth India 2023 Go Server online!!",
		})
	})

	r.POST("/ask-ai", func(c *gin.Context) {
		var reqBody ReqBody

		// Bind the JSON body to the struct
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Use the 'msg' from the request
		msg := reqBody.Msg

		// Your logic here...
		answer := FetchOpenAI(msg, openAIConstants)

		c.JSON(http.StatusOK, gin.H{"answer": answer})
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
