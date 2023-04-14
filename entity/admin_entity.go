package entity

type AdminEntity struct {
	Id       int64  `gorm:"primary_key;autoIncrement" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Role     string `gorm:"column:role" json:"role"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (AdminEntity) TableName() string {
	return "admin"
}
