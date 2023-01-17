package main

import (
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const KeyRequestID = "requestId"

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 请求新进入这里
	r.Use(func(context *gin.Context) {
		s := time.Now()
		context.Next()

		// path , log , latency ,response  ,code
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	},
		func(context *gin.Context) {
			context.Set("requestId", rand.Int())
			context.Next()
		},
	)

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestId, exists := c.Get("requestId"); exists {
			h[KeyRequestID] = requestId
		}
		c.JSON(http.StatusOK, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
