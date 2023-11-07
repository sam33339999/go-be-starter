package models

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Age       uint8  `json:"age"`
	// Birthday  string `json:"birthday"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}

func NewUser() User {
	return User{}
}
