package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
	"time"
)

var todos = Todos{
	Todo{
		Id:        100,
		Name:      "Write GO lang REST api",
		Completed: true,
		Due:       time.Date(2016, time.November, 2, 11, 0, 0, 0, time.UTC),
	},
	Todo{
		Id:        101,
		Name:      "Buy Milk",
		Completed: false,
		Due:       time.Date(2016, time.November, 2, 23, 0, 0, 0, time.UTC),
	},
	Todo{
		Id:        102,
		Name:      "Talk to Justin and Alex",
		Completed: false,
		Due:       time.Date(2016, time.November, 4, 23, 0, 0, 0, time.UTC),
	},
}

// curl -i http://localhost:8080/api/v1/todos
func GetTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

// curl -i http://localhost:8080/api/v1/todos/100
func GetTodo(context *gin.Context) {
	id := context.Params.ByName("todoid")
	todoId, _ := strconv.ParseInt(id, 10, 64)

	var todo Todo

	for _, elem := range todos {
		if elem.Id == todoId {
			todo = elem
			break
		}
	}

	if IsEmpty(todo) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo item was not found"})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// curl -i -X POST -H "Content-Type: application/json" -d "{}" http://localhost:8080/api/v1/todos
func CreateTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "CreateTodo",
	})
}

// curl -i -X PUT -H "Content-Type: application/json" -d "{}" http://localhost:8080/api/v1/todos/100
func UpdateTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "UpdateTodo",
	})
}

// curl -i -X DELETE http://localhost:8080/api/v1/todos/100
func DeleteTodo(context *gin.Context) {
	id := context.Params.ByName("todoid")
	todoId, _ := strconv.ParseInt(id, 10, 64)
	found := false

	for index, elem := range todos {
		if elem.Id == todoId {
			todos = append(todos[:index], todos[index+1:]...)
			found = true
			break
		}
	}

	if !found {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo item was not found"})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Todo item was deleted"})
	}
}
