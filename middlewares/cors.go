package middlewares

import (
  "fmt"
  "net/http"
  "gopkg.in/gin-gonic/gin.v1"
)

// taken from https://github.com/Massad/gin-boilerplate/blob/master/main.go#L13-L30
// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
    c.Writer.Header().Set("Access-Control-Max-Age", "86400")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
    c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

    if c.Request.Method == "OPTIONS" {
      fmt.Println("OPTIONS")
      c.AbortWithStatus(http.StatusOK)
    } else {
      c.Next()
    }
  }
}