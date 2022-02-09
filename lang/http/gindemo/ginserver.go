package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(
		func(c *gin.Context) {
			c.Set(keyRequestId, rand.Int())

			c.Next()
		},
		func(c *gin.Context) {

			s := time.Now()

			c.Next()
			fields := []zapcore.Field{
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.Duration("elapsed", time.Now().Sub(s)),
			}
			if rid, exists := c.Get(keyRequestId); exists {
				fields = append(fields, zap.Int(keyRequestId, rid.(int)))
			}

			logger.Info("incoming request", fields...)
		})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
