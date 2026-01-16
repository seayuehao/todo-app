package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/todo-app/internal/dto"
	"github.com/todo-app/internal/service"
	"github.com/todo-app/pkg/base"
	"log"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	svc *service.TodoService
}

func NewTodoHandler(svc *service.TodoService) *TodoHandler {
	return &TodoHandler{
		svc: svc,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	userID, err := base.MustAuth(c)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	var req dto.AddTodoReq
	if err := c.BindJSON(&req); err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	todo, err := h.svc.Add(c.Request.Context(), &req, userID)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) ListTodos(c *gin.Context) {
	userID, err := base.MustAuth(c)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	todos, err := h.svc.List(c.Request.Context(), userID)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userID, err := base.MustAuth(c)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	//var baseIdReq dto.BaseIdReq
	//if err  = c.BindUri(&baseIdReq); err != nil {
	//	return err
	//}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}

	err = h.svc.Delete(c.Request.Context(), id, userID)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TodoHandler) CompleteTodo(c *gin.Context) {
	userID, err := base.MustAuth(c)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}

	data, err := h.svc.Complete(c.Request.Context(), id, userID)
	if err != nil {
		log.Printf("error occurred: %v", err)
		base.StdErr(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
