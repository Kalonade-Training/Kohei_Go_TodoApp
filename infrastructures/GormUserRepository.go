// infrastructure/database/repository/user_repository.go
package infrastructures 

import (
	"fmt"

	"goTodoApp/domain/entities"
	"goTodoApp/domain/repositories"
	"goTodoApp/infrastructures/mapper"
	value_object "goTodoApp/domain/value-object"	
	"goTodoApp/infrastructures/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &GormUserRepository{db: db}
}

// Save はユーザーを保存（新規 or 更新）する
func (r *GormUserRepository) Save(user *entities.User) (*entities.User, error) {
	userModel := mapper.EntityToUserModel(*user)

	if err := r.db.Save(&userModel).Error; err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	savedEntity, err := mapper.ModelToUserEntity(userModel)
	if err != nil {
		return nil, err
	}
	return savedEntity, nil
}

// FindByID はIDでユーザーを取得する
func (r *GormUserRepository) FindByUsername(username string) (*entities.User, error) {
    var userModel model.User
		usernameVo, err := value_object.NewUsername(username)
		// NewUsernameは戻り値にerrorがあるのでエラーチェックが必要NewUsername(username) 
    if err != nil {
        return nil, err
    }
    err = r.db.Where("username = ?", usernameVo.Value()).First(&userModel).Error
    if err != nil {
      return nil, err
    }

    // model -> entity 変換
    userEntity, err := mapper.ModelToUserEntity(userModel)
		if err != nil {
			return nil, err
		}

    return userEntity, nil
}