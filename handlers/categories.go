package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mrityunjayvashisth/mini-api-demo/generated/categories"
)

// CategoryHandler implements the ServerInterface for category operations
type CategoryHandler struct {
	categories map[string]categories.Category
	mu         sync.RWMutex
}

// NewCategoryHandler creates a new handler for category operations
func NewCategoryHandler() *CategoryHandler {
	// Create sample data
	now := time.Now()
	// tomorrow := now.Add(24 * time.Hour)

	initialCategories := map[string]categories.Category{
		"category-1": {
			Id:          "category-1",
			Name:        "Work",
			Description: stringPtr("Work-related tasks"),
			Color:       stringPtr("#FF5733"),
			CreatedAt:   &now,
		},
		"category-2": {
			Id:          "category-2",
			Name:        "Personal",
			Description: stringPtr("Personal errands and tasks"),
			Color:       stringPtr("#33FF57"),
			CreatedAt:   &now,
		},
	}

	return &CategoryHandler{
		categories: initialCategories,
	}
}

// ListCategories handles GET /categories
func (h *CategoryHandler) ListCategories(ctx echo.Context, params categories.ListCategoriesParams) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// Prepare result slice
	result := make([]categories.Category, 0, len(h.categories))

	for _, category := range h.categories {
		result = append(result, category)

		// Apply limit
		if params.Limit != nil && len(result) >= *params.Limit {
			break
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

// CreateCategory handles POST /categories
func (h *CategoryHandler) CreateCategory(ctx echo.Context) error {
	var input categories.CategoryInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	// Generate a new ID
	id := uuid.New().String()
	now := time.Now()

	// Create the category
	category := categories.Category{
		Id:          id,
		Name:        input.Name,
		Description: input.Description,
		Color:       input.Color,
		CreatedAt:   &now,
	}

	// Save it
	h.categories[id] = category

	return ctx.JSON(http.StatusCreated, category)
}

// GetCategoryById handles GET /categories/{categoryId}
func (h *CategoryHandler) GetCategoryById(ctx echo.Context, categoryId string) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	category, exists := h.categories[categoryId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	return ctx.JSON(http.StatusOK, category)
}

// UpdateCategory handles PUT /categories/{categoryId}
func (h *CategoryHandler) UpdateCategory(ctx echo.Context, categoryId string) error {
	var input categories.CategoryInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	category, exists := h.categories[categoryId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	// Update fields
	category.Name = input.Name
	category.Description = input.Description
	category.Color = input.Color

	// Save updated category
	h.categories[categoryId] = category

	return ctx.JSON(http.StatusOK, category)
}

// DeleteCategory handles DELETE /categories/{categoryId}
func (h *CategoryHandler) DeleteCategory(ctx echo.Context, categoryId string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.categories[categoryId]; !exists {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}

	delete(h.categories, categoryId)

	return ctx.NoContent(http.StatusNoContent)
}
