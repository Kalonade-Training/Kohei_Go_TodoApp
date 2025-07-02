package infrastructures

import (
	"goTodoApp/domain/repositories"
	"goTodoApp/domain/entities"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository はリポジトリのインスタンスを生成
func NewGormUserRepository(db *gorm.DB) repositories.IUserRepository{
	return &GormUserRepository{db: db}
}

//ユーザー情報を保存
func (r *GormUserRepository) Save(user *entities.User) error{
	return r.db.Save(user).Error
}

//ユーザーを検索
func (r *GormUserRepository) FindByUsername(username string) (*entities.User, error){
	var user entities.User
	//PKではないからWhere.Firstにする必要がある.
	if err := r.db.Where("username = ? ", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}