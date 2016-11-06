package models

import (
	"time"
	"errors"
)

// Todo ...
type Todo struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	UpdatedAt *time.Time `json:updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

// This is just for the example ...
type Todos []Todo

var todos = Todos{
	Todo{
		Id:        100,
		Text:      "Write GO lang REST api",
		Completed: true,
		Due:       time.Date(2016, time.November, 2, 11, 0, 0, 0, time.UTC),
		UpdatedAt: nil,
		CreatedAt: time.Date(2016, time.January, 1, 9, 30, 0, 0, time.UTC),
	},
	Todo{
		Id:        101,
		Text:      "Buy Milk",
		Completed: false,
		Due:       time.Date(2016, time.November, 2, 23, 0, 0, 0, time.UTC),
		UpdatedAt: nil,
		CreatedAt: time.Date(2016, time.January, 1, 9, 30, 0, 0, time.UTC),
	},
	Todo{
		Id:        102,
		Text:      "Talk to Justin and Alex",
		Completed: false,
		Due:       time.Date(2016, time.November, 4, 23, 0, 0, 0, time.UTC),
		UpdatedAt: nil,
		CreatedAt: time.Date(2016, time.January, 1, 9, 30, 0, 0, time.UTC),
	},
}

var maxId int64 = 102

// TypeModel ...
type TodoModel struct{}

// One ...
func (m TodoModel) One(todoId int64) (todo Todo, err error) {
	todoIndex, todo := getTodoById(todoId, todos)

	if todoIndex == -1 {
		return todo, errors.New("Todo not found")
	} else {
		return todo, nil
	}
}

// All ...
func (m TodoModel) All() (data Todos, err error) {
	return todos, nil
}

// Create ...
func (m TodoModel) Create(todo Todo) (todoId int64, err error) {
	maxId = maxId + 1

	newTodo := Todo{
		Id:        maxId,
		Text:      todo.Text,
		Completed: todo.Completed,
		Due:       todo.Due,
		UpdatedAt: nil,
		CreatedAt: time.Now().UTC(),
	}

	todos = append(todos, newTodo)

	return maxId, err
}

// Update ...
func (m TodoModel) Update(todoId int64, todo Todo) (err error) {
	todoIndex, oldTodo := getTodoById(todoId, todos)

	if todoIndex == -1 {
		return errors.New("Todo not found")
	}

	// updating fields
	if !IsEmpty(todo.Text) {
		oldTodo.Text = todo.Text
	}

	if !IsEmpty(todo.Completed) {
		oldTodo.Completed = todo.Completed
	}

	if !IsEmpty(todo.Due) {
		oldTodo.Due = todo.Due
	}

	nowTime := time.Now().UTC()
	oldTodo.UpdatedAt = &nowTime

	todos[todoIndex] = oldTodo

	return err
}

// Delete ...
func (m TodoModel) Delete(todoId int64) (err error) {
	todoIndex, _ := getTodoById(todoId, todos)

	if todoIndex == -1 {
		return errors.New("Todo not found")
	}

	todos = append(todos[:todoIndex], todos[todoIndex+1:]...)

	return err
}

func getTodoById(todoId int64, todos Todos) (int, Todo) {
	for index, elem := range todos {
		if elem.Id == todoId {
			return index, elem
		}
	}

	return -1, Todo{}
}
