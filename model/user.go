package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "user"
}
