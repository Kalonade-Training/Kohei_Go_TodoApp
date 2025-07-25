package infrastructures

import (
	"errors"
	"time"
	"fmt"

	"gorm.io/gorm"
	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	value_object "goTodoApp/domain/value-object"
	"goTodoApp/infrastructures/mapper"
	"goTodoApp/infrastructures/model"
)

type GormTodoRepository struct {
	db *gorm.DB
}

// GormTodoRepository を作成
func NewGormTodoRepository(db *gorm.DB) repositories.ITodoRepository {
	return &GormTodoRepository{db: db}
}

// Todoを保存
func (r *GormTodoRepository) Save(todo *entities.Todo) (*entities.Todo, error) {
	if todo.ID().Value() == "" {
		return nil, errors.New("todo ID is required")
	}

	// 値が初期なら現在時刻をセット
	if todo.CreatedAt().IsZero() {
		todo.SetCreatedAt(time.Now())
	}
	todo.SetUpdatedAt(time.Now())

	todoModel := mapper.EntityToTodoModel(*todo)

	err := r.db.Save(&todoModel).Error
	if err != nil {
		return nil, err
	}

	savedEntity, err := mapper.ModelToEntity(todoModel)
	if err != nil {
		return nil, err
	}
	return savedEntity, nil
}

// 指定されたIDのTodoを取得
func (r *GormTodoRepository) FindTodoByID(
	todoID value_object.TodoID,
	userID value_object.UserID,
) (*entities.Todo, error) {
	var todoModel model.Todo
	if err := r.db.Where(
		"id = ? AND user_id = ?",
		todoID.Value(), 
		userID.Value(),
	).First(&todoModel).Error; err != nil {
		return nil, err
	}
	todo, err := mapper.ModelToEntity(todoModel)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// 全検索+filter検索
func (r *GormTodoRepository) FindByUserIDWithFilters(userID value_object.UserID, filters *repositories.TodoFilters) ([]*entities.Todo, error) {
	var models []model.Todo
	query := r.db.Model(&model.Todo{}).Where("user_id = ?", userID.Value())

	if filters.Title != nil && *filters.Title != "" {
		query = query.Where("title LIKE ?", "%"+ *filters.Title +"%")
	}
	if filters.Description != nil && *filters.Description != "" {
    query = query.Where("description LIKE ?", "%"+ *filters.Description +"%")
	}
	if filters.DueDateFrom != nil {
			query = query.Where("due_date >= ?", filters.DueDateFrom)
	}
	if filters.DueDateTo != nil {
			query = query.Where("due_date <= ?", filters.DueDateTo)
	}
	if filters.Status != nil {
		switch *filters.Status {
			case "completed":
				query = query.Where("status = ?", "completed")
			case "in_progress":
				query = query.Where("status = ?", "in_progress")
		}
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find todos: %w", err)
	}
	todos := make([]*entities.Todo, 0, len(models))
	for _, m := range models {
		entity, err := mapper.ModelToEntity(m)
		if err != nil {
			return nil, err
		}
		todos = append(todos, entity)
	}
	return todos, nil
}

// Delete は指定されたIDのTodoを削除します
func (r *GormTodoRepository) Delete(todoID value_object.TodoID, userID value_object.UserID, ) error {
	result := r.db.Where("id = ? AND user_id = ?",todoID.Value(), userID.Value()).Delete(&model.Todo{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("todo not found or unauthorized")
    }
    return nil
}

// Duplicate は指定した Todo を複製する
func (r *GormTodoRepository) Duplicate(todoid value_object.TodoID, userid value_object.UserID) (*entities.Todo, error) {
	var originalModel model.Todo
// ユーザーIDも条件に加えることで、自分のTodoだけを複製できるようにする
	if err := r.db.First(&originalModel, "id = ? AND user_id = ?", todoid.Value(), userid.Value()).Error; err != nil {
		return nil, err
	}

	// モデル → エンティティに変換
	originalEntity, err := mapper.ModelToEntity(originalModel)
	if err != nil {
		return nil, err
	}

	// エンティティを複製（IDなどは新しく振られる想定）
	duplicatedEntity, err := originalEntity.Duplicate()
	if err != nil {
    return nil, err
	}
	// エンティティ → モデルに変換
	duplicatedModel := mapper.EntityToTodoModel(*duplicatedEntity)

	// DBに保存
	if err := r.db.Create(&duplicatedModel).Error; err != nil {
		return nil, err
	}

	return duplicatedEntity, nil
}


func (r *GormTodoRepository) Update(todo *entities.Todo) (*entities.Todo, error) {
	// エンティティ → モデルへ変換
	todoModel := mapper.EntityToTodoModel(*todo)

	// GORMで更新
	if err := r.db.Save(&todoModel).Error; err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	// DBから最新データを取得
	var updatedModel model.Todo
	if err := r.db.First(&updatedModel, "id = ?", todoModel.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated todo: %w", err)
	}

	// モデル → エンティティに変換して返す
	updatedEntity, err := mapper.ModelToEntity(updatedModel)
	if err != nil {
		return nil, err
	}
	return updatedEntity, nil
}