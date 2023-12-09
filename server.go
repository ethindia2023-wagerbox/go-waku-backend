package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"
const INIT_MSG = "Eth India 2023 Go Server online!!"

func InitServer(openAIConstants OpenAIConstants) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": INIT_MSG,
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

	return r
}

func StartServer(r *gin.Engine) {
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(PORT)
}
