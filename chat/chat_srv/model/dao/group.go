package dao

type Group struct {
	Base
	UserID uint32 `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
}

func (m *Group) TableName() string {
	return "chat_group"
}
