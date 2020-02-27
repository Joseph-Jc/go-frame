package model

// 自动迁移
func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Student{})
}
