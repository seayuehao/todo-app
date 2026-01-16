package service

import "github.com/todo-app/internal/db"

type ServiceGroup struct {
	UserService *UserService
	TodoService *TodoService
}

// NewServiceGroup 初始化服务组
func NewServiceGroup() *ServiceGroup {
	return initServiceGroup()
}

func initServiceGroup() *ServiceGroup {
	userSvc := NewUserService(db.DB)
	todoSvc := NewTodoService(db.DB)

	svcgrp := ServiceGroup{
		UserService: userSvc,
		TodoService: todoSvc,
	}
	return &svcgrp
}
