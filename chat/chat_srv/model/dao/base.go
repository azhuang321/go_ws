package dao

type Base struct {
	ID        uint `json:"id" gorm:"column:id"`
	CreatedAt uint `json:"created_at" gorm:"column:created_at"`
	UpdatedAt uint `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt uint `json:"deleted_at" gorm:"column:deleted_at"`
}
