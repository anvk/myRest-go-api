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
		Text:      "Write GO lang REST api",
		Completed: true,
		Due:       time.Date(2016, time.November, 2, 11, 0, 0, 0, time.UTC),
	},
	Todo{
		Id:        101,
		Text:      "Buy Milk",
		Completed: false,
		Due:       time.Date(2016, time.November, 2, 23, 0, 0, 0, time.UTC),
	},
	Todo{
		Id:        102,
		Text:      "Talk to Justin and Alex",
		Completed: false,
		Due:       time.Date(2016, time.November, 4, 23, 0, 0, 0, time.UTC),
	},
}

var maxId int64 = 102

// curl -i http://localhost:8080/api/v1/todos
func GetTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

// curl -i http://localhost:8080/api/v1/todos/100
func GetTodo(context *gin.Context) {
	id := context.Params.ByName("todoid")
	todoId, _ := strconv.ParseInt(id, 10, 64)

	_, todo := getTodoById(todoId, todos)

	if IsEmpty(todo) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo item was not found"})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// curl -i -X POST -H "Content-Type: application/json" -d "{\"Text\": \"New Todo Item\", \"Completed\": false, \"Due\":\"2016-11-01T12:30:00Z\"}" http://localhost:8080/api/v1/todos
func CreateTodo(context *gin.Context) {
	var json Todo

	context.Bind(&json)

	if IsEmpty(json.Text) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Text field cannot be empty"})
		return
	}

	maxId = maxId + 1

	todo := Todo{
		Id:        maxId,
		Text:      json.Text,
		Completed: json.Completed,
		Due:       json.Due,
	}

	todos = append(todos, todo)

	context.JSON(http.StatusOK, gin.H{"message": "Todo item was added"})
}

// curl -i -X PUT -H "Content-Type: application/json" -d "{\"Text\": \"Todo Item Changed\"}" http://localhost:8080/api/v1/todos/100
func UpdateTodo(context *gin.Context) {
	id := context.Params.ByName("todoid")
	todoId, _ := strconv.ParseInt(id, 10, 64)
	var json Todo

	todoIndex, todo := getTodoById(todoId, todos)

	if IsEmpty(todo) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo item was not found"})
		return
	}

	context.Bind(&json)

	// updating fields
	if !IsEmpty(json.Text) {
		todo.Text = json.Text
	}

	if !IsEmpty(json.Completed) {
		todo.Completed = json.Completed
	}

	if !IsEmpty(json.Due) {
		todo.Due = json.Due
	}

	todos[todoIndex] = todo

	context.JSON(http.StatusOK, gin.H{"message": "Todo item was updated"})
}

// curl -i -X DELETE http://localhost:8080/api/v1/todos/100
func DeleteTodo(context *gin.Context) {
	id := context.Params.ByName("todoid")
	todoId, _ := strconv.ParseInt(id, 10, 64)

	todoIndex, _ := getTodoById(todoId, todos)

	if todoIndex == -1 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo item was not found"})
		return
	}

	todos = append(todos[:todoIndex], todos[todoIndex+1:]...)
	context.JSON(http.StatusOK, gin.H{"message": "Todo item was deleted"})
}

func getTodoById(todoId int64, todos Todos) (int, Todo) {
	for index, elem := range todos {
		if elem.Id == todoId {
			return index, elem
		}
	}

	return -1, Todo{}
}
