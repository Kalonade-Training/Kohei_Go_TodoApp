package controllers

import (
	"net/http"
	"time"
	"io"
	"fmt"
	"goTodoApp/usecases"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	createTodo   *usecases.CreateTodo
	findTodoByID *usecases.FindTodoByID
	findAllTodos *usecases.FindAllTodos
	updateTodo   *usecases.UpdateTodo
	deleteTodo   *usecases.DeleteTodo
	duplicateTodo *usecases.DuplicateTodo
}

func NewTodoController(
	create *usecases.CreateTodo,
	findByID *usecases.FindTodoByID,
	findAll *usecases.FindAllTodos,
	update *usecases.UpdateTodo,
	deleteUC *usecases.DeleteTodo,
	duplicate *usecases.DuplicateTodo,
) *TodoController {
	return &TodoController{
		createTodo:    create,
		findTodoByID:  findByID,
		findAllTodos:  findAll,
		updateTodo:    update,
		deleteTodo:    deleteUC,
		duplicateTodo: duplicate,
	}
}

// Create ハンドラー例
func (tc *TodoController) Create(c *gin.Context) {
	var req struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description"`
		DueDate     *string `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dueDate *time.Time
	if req.DueDate != nil {
		t, err := time.Parse("2006-01-02", *req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format. Use YYYY-MM-DD"})
			return
		}
		dueDate = &t
	}

	todo, err := tc.createTodo.Execute(req.Title, req.Description, dueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) FindAll(c *gin.Context) {
	todos, err := tc.findAllTodos.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) FindByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := tc.findTodoByID.Execute(id)
	if err != nil {
		if err.Error() == "todo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) Update(c *gin.Context) {
	id := c.Param("id")

	var fields map[string]interface{}
	err := c.ShouldBindJSON(&fields)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	fmt.Printf("Update request body fields: %+v\n", fields)//デバック

	if len(fields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
		return
	}

	updatedTodo, err := tc.updateTodo.Execute(id, fields)
	if err != nil {
		if err.Error() == "todo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func (tc *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := tc.deleteTodo.Execute(id); err != nil {
		if err.Error() == "todo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil) // 成功時は204 No Content
}

func (tc *TodoController) Duplicate(c *gin.Context) {
	id := c.Param("id")
	
	if id == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "todo ID is required"})
    return
	}
	duplicatedTodo, err := tc.duplicateTodo.Execute(id)
	if err != nil {
		if err.Error() == "todo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, duplicatedTodo) // 成功時は201 Created
}