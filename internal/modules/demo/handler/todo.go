package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/modules/demo/service"
	"github.com/rei0721/go-scaffold/pkg/logger"
	"github.com/rei0721/go-scaffold/types/result"
)

type TodoHandler struct {
	service service.TodoService
	logger  logger.Logger
}

type createTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type updateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func NewTodoHandler(service service.TodoService, logger logger.Logger) *TodoHandler {
	return &TodoHandler{service: service, logger: logger}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}

	todo, err := h.service.Create(c.Request.Context(), service.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result.Success(todo))
}

func (h *TodoHandler) List(c *gin.Context) {
	todos, err := h.service.List(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, todos)
}

func (h *TodoHandler) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	todo, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, todo)
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var req updateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}

	todo, err := h.service.Update(c.Request.Context(), id, service.UpdateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, todo)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, gin.H{"deleted": true})
}

func parseID(c *gin.Context) (uint, bool) {
	value, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || value == 0 {
		result.BadRequest(c, "invalid id")
		return 0, false
	}
	return uint(value), true
}

func (h *TodoHandler) writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrTodoTitleRequired):
		result.BadRequest(c, err.Error())
	case errors.Is(err, service.ErrTodoNotFound):
		result.NotFound(c, err.Error())
	default:
		if h.logger != nil {
			h.logger.Error("todo request failed", "error", err)
		}
		result.InternalError(c, "internal server error")
	}
}
