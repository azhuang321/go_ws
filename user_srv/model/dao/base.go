package dao

type Base struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
	DeletedAt uint `json:"deleted_at"`
}
