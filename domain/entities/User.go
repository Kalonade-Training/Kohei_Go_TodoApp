package entities
import "time"
type User struct {
	ID string `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	// bcryptハッシュ保存 
	Password string `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`                    // 作成日時を自動セット
  UpdatedAt   time.Time  `gorm:"autoUpdateTime"`                    // 更新日時を自動セット
}