package repository

import (
	"github.com/google/uuid"
	"github.com/todo-app/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(entity *models.Todo) error {
	return r.db.Create(entity).Error
}

func (r *TodoRepository) GetByID(id int) (*models.Todo, error) {
	var entity models.Todo
	err := r.db.Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *TodoRepository) Delete(id int) error {
	var entity models.Todo
	return r.db.Where("id = ?", id).Delete(&entity).Error
}

func (r *TodoRepository) Update(entity *models.Todo) error {
	return r.db.Save(entity).Error
}

func (r *TodoRepository) FindAllByUserId(userId uuid.UUID) ([]models.Todo, error) {
	var arr []models.Todo
	err := r.db.Where("user_id = ?", userId).Find(&arr).Error
	if err != nil {
		return nil, err
	}
	return arr, nil
}
