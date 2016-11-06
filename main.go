package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"

	"github.com/anvk/myRest-go-api/controllers"
	"github.com/anvk/myRest-go-api/middlewares"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	v1 := router.Group("api/v1")
	{
		// curl -i http://localhost:8080/ping
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Message": "pong"})
		})

		/*** START Todos ***/
		todo := new(controllers.TodoController)

		// curl -i http://localhost:8080/api/v1/todos
		v1.GET("/todos", todo.All)
		// curl -i http://localhost:8080/api/v1/todos/100
		v1.GET("/todos/:todoid", todo.One)
		// curl -i -X POST -H "Content-Type: application/json" -d "{\"Text\": \"New Todo Item\", \"Completed\": false, \"Due\":\"2016-11-01T12:30:00Z\"}" http://localhost:8080/api/v1/todos
		v1.POST("/todos", todo.Create)
		// curl -i -X PUT -H "Content-Type: application/json" -d "{\"Text\": \"Todo Item Changed\"}" http://localhost:8080/api/v1/todos/100
		v1.PUT("/todos/:todoid", todo.Update)
		// curl -i -X DELETE http://localhost:8080/api/v1/todos/100
		v1.DELETE("/todos/:todoid", todo.Delete)
	}

	router.Run(":8080") // listen and server on 0.0.0.0:8080
}
