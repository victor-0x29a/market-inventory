package database

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:128"`
	Email string `gorm:"uniqueIndex"`
	Role  string `gorm:"type:enum('admin','student','teacher');not null"`
}
