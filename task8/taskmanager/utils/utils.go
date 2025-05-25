package utils

import "github.com/gin-gonic/gin"

func HTTP(handler func(*gin.Context)) gin.HandlerFunc {
    return func(c *gin.Context) {
        handler(c)
    }
}