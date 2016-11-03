package main

import "gopkg.in/gin-gonic/gin.v1"

func Options(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	context.Next()
}
