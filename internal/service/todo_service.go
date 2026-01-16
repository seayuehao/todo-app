package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/todo-app/internal/dto"
	"github.com/todo-app/internal/models"
	"github.com/todo-app/internal/repository"
	"gorm.io/gorm"
	"time"
)

type TodoService struct {
	db   *gorm.DB
	repo *repository.TodoRepository
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		db:   db,
		repo: repository.NewTodoRepository(db),
	}
}

func (s *TodoService) Add(ctx context.Context, req *dto.AddTodoReq, userId uuid.UUID) (*models.Todo, error) {
	todo := models.Todo{
		UserID:    userId,
		CreatedAt: time.Now(),
		Title:     req.Title,
	}

	err := s.repo.Create(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, err
}

func (s *TodoService) List(ctx context.Context, userId uuid.UUID) ([]models.Todo, error) {
	arr, err := s.repo.FindAllByUserId(userId)
	return arr, err
}

func (s *TodoService) Delete(ctx context.Context, id int, userId uuid.UUID) error {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if res.UserID != userId {
		return errors.New("resource not found")
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) Complete(ctx context.Context, id int, userId uuid.UUID) (*models.Todo, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if res.UserID != userId {
		return nil, errors.New("resource not found")
	}

	now := time.Now()
	res.Completed = true
	res.UpdatedAt = &now
	err = s.repo.Update(res)
	return res, err
}
