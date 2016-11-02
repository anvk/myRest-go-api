package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
    router := gin.Default()

    router.Use(Cors())

    v1 := router.Group("api/v1")
    {
        v1.GET("/todos", GetTodos)
        v1.GET("/todos/:todoid", GetTodo)
        v1.POST("/todos", CreateTodo)
        v1.PUT("/todos/:todoid", UpdateTodo)
        v1.DELETE("/todos/:todoid", DeleteTodo)
        v1.OPTIONS("/todos", Options)         // POST
        v1.OPTIONS("/todos/:todoid", Options) // PUT, DELETE
    }

    router.GET("/ping", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.Run() // listen and server on 0.0.0.0:8080
}