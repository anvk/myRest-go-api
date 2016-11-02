package main

import "gopkg.in/gin-gonic/gin.v1"

func Cors() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
        context.Next()
    }
}