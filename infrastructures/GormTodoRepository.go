package infrastructures

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
)

type GormTodoRepository struct {
	db *gorm.DB
}

// NewGormTodoRepository は GormTodoRepository を作成します
func NewGormTodoRepository(db *gorm.DB) repositories.ITodoRepository {
	return &GormTodoRepository{db: db}
}

// Save はTodoを保存します
func (r *GormTodoRepository) Save(todo *entities.Todo) error {
	if todo.ID == "" {
		return errors.New("todo ID is required")
	}
	
	// CreatedAtは新規作成時のみセットするなど設計次第で調整
	if todo.CreatedAt.IsZero() {
		todo.CreatedAt = time.Now()
	}
	todo.UpdatedAt = time.Now()
	return r.db.Save(todo).Error
}

// FindByID は指定されたIDのTodoを取得します
func (r *GormTodoRepository) FindByID(id string) (*entities.Todo, error) {
	var todo entities.Todo
	if err := r.db.First(&todo, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

// FindAll は全てのTodoを取得します
func (r *GormTodoRepository) FindAll() ([]*entities.Todo, error) {
	var todos []*entities.Todo
	if err := r.db.Order("created_at DESC").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Delete は指定されたIDのTodoを削除します
func (r *GormTodoRepository) Delete(id string) error {
	result := r.db.Delete(&entities.Todo{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return result.Error
}

// Duplicate は指定されたIDのTodoを複製します
func (r *GormTodoRepository) Duplicate(id string) (*entities.Todo, error) {
	original, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}

	newTodo := original.Duplicate()
	err = r.Save(newTodo)
	return newTodo, err
}


func (r *GormTodoRepository) Update(id string, fields map[string]interface{}) error {
	// 更新タイムスタンプを追加
    fields["updated_at"] = time.Now()
    if err := r.db.Model(&entities.Todo{}).Where("id = ?", id).Updates(fields).Error; err != nil {
        return err
    }
    return nil
}