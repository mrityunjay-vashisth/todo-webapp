package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mrityunjay-vashisth/mini-api-demo/generated/todos"
)

// TodoHandler implements the ServerInterface for todo operations
type TodoHandler struct {
	todos map[string]todos.Todo
	mu    sync.RWMutex
}

// NewTodoHandler creates a new handler for todo operations
func NewTodoHandler() *TodoHandler {
	// Create sample data
	now := time.Now()
	tomorrow := time.Now().Add(24 * time.Hour)

	initialTodos := map[string]todos.Todo{
		"todo-1": {
			Id:          "todo-1",
			Title:       "Complete project",
			Description: stringPtr("Finish the API project by tomorrow"),
			Completed:   boolPtr(false),
			CategoryId:  stringPtr("category-1"),
			CreatedAt:   &now,
			DueDate:     &tomorrow,
		},
		"todo-2": {
			Id:          "todo-2",
			Title:       "Buy groceries",
			Description: stringPtr("Milk, eggs, bread"),
			Completed:   boolPtr(true),
			CategoryId:  stringPtr("category-2"),
			CreatedAt:   &now,
		},
	}

	return &TodoHandler{
		todos: initialTodos,
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

// ListTodos handles GET /todos
func (h *TodoHandler) ListTodos(ctx echo.Context, params todos.ListTodosParams) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Prepare result slice
	result := make([]todos.Todo, 0)

	// Apply filters
	for _, todo := range h.todos {
		// Filter by completion status if specified
		if params.Completed != nil && *params.Completed != *todo.Completed {
			continue
		}

		result = append(result, todo)

		// Apply limit
		if params.Limit != nil && len(result) >= *params.Limit {
			break
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

// CreateTodo handles POST /todos
func (h *TodoHandler) CreateTodo(ctx echo.Context) error {
	var input todos.TodoInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	// Generate a new ID
	id := uuid.New().String()
	now := time.Now()

	// Create the todo
	todo := todos.Todo{
		Id:          id,
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
		CategoryId:  input.CategoryId,
		CreatedAt:   &now,
		DueDate:     input.DueDate,
	}

	// Save it
	h.todos[id] = todo

	return ctx.JSON(http.StatusCreated, todo)
}

// GetTodoById handles GET /todos/{todoId}
func (h *TodoHandler) GetTodoById(ctx echo.Context, todoId string) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	todo, exists := h.todos[todoId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	return ctx.JSON(http.StatusOK, todo)
}

// UpdateTodo handles PUT /todos/{todoId}
func (h *TodoHandler) UpdateTodo(ctx echo.Context, todoId string) error {
	var input todos.TodoInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	todo, exists := h.todos[todoId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	// Update fields
	todo.Title = input.Title
	todo.Description = input.Description
	todo.Completed = input.Completed
	todo.CategoryId = input.CategoryId
	todo.DueDate = input.DueDate

	// Save updated todo
	h.todos[todoId] = todo

	return ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE /todos/{todoId}
func (h *TodoHandler) DeleteTodo(ctx echo.Context, todoId string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.todos[todoId]; !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	delete(h.todos, todoId)

	return ctx.NoContent(http.StatusNoContent)
}
