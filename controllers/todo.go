package controllers

import (
  "strconv"
  "net/http"
  "github.com/anvk/myRest-go-api/models"

  "gopkg.in/gin-gonic/gin.v1"
)

// ArticleController ...
type TodoController struct{}

var todoModel = new(models.TodoModel)

// One ...
func (ctrl TodoController) One(c *gin.Context) {
  id := c.Params.ByName("todoid")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {
    data, err := todoModel.One(id)
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"Message": "Todo item not found", "error": err.Error()})
      c.Abort()
      return
    }

    c.JSON(http.StatusOK, data)
  } else {
    c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
  }
}

// All ...
func (ctrl TodoController) All(c *gin.Context) {
  data, err := todoModel.All()

  if err != nil {
    c.JSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get todos", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(http.StatusOK, data)
}

// Create ...
func (ctrl TodoController) Create(c *gin.Context) {
  var todo models.Todo

  if c.BindJSON(&todo) != nil {
    c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": todo})
    c.Abort()
    return
  }

  todoId, err := todoModel.Create(todo)

  if todoId > 0 && err != nil {
    c.JSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Todo item created", "id": todoId})
}

// Update ...
func (ctrl TodoController) Update(c *gin.Context) {
  id := c.Params.ByName("todoid")
  var todo models.Todo

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {
    if c.BindJSON(&todo) != nil {
      c.JSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form", "form": todo})
      c.Abort()
      return
    }

    err := todoModel.Update(id, todo)
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"Message": "Todo item not found", "error": err.Error()})
      c.Abort()
      return
    }

    c.JSON(http.StatusOK, gin.H{"Message": "Todo item updated"})
  } else {
    c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
  }
}

// Delete ...
func (ctrl TodoController) Delete(c *gin.Context) {
  id := c.Params.ByName("todoid")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {
    err := todoModel.Delete(id)

    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"Message": "Todo item not found", "error": err.Error()})
      c.Abort()
      return
    }

    c.JSON(http.StatusOK, gin.H{"Message": "Todo item deleted"})
  } else {
    c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
  }
}