package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/atirach-in/management-system.git/configs"
	"github.com/atirach-in/management-system.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Using .env
	if err := configs.InitDotEnv(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	PORT := os.Getenv("PORT")

	// Enable log's color
	gin.ForceConsoleColor()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("\n\n\n%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/health/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Service API Project management is running port " + PORT,
		})
	})

	//Connect DB
	configs.TestConnectDB()

	routes.InitRoute(r)

	fmt.Println("\n\nThis server running port: " + PORT)
	r.Run(PORT) // listen and serve on 0.0.0.0:8080
}
