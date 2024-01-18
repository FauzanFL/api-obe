package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserDosen struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	KodeDosen string `json:"kode_dosen"`
	Nama      string `json:"nama"`
}

func (User) TableName() string {
	return "user"
}
