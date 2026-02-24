package user

import "github.com/jinzhu/gorm"

// Player 定义数据库模型
type Player struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Name       string `gorm:"size:255;not null"`
	Hero       string `gorm:"size:255"`
	MMR        int
}

// 简单的数据库操作函数
func CreateRecord(db *gorm.DB) error {
	record := map[string]interface{}{
		"name": "test",
		"age":  25,
	}
	return db.Table("users").Create(record).Error
}

// 被mock的函数
func GetUserID() uint64 {
	return 123
}
