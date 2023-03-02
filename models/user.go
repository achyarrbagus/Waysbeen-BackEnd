package models

import "time"

// user modal struct
type User struct {
	ID        int             `json:"id"`
	Name      string          `json:"name" gorm:"type: varchar(225)"`
	Email     string          `json:"email" gorm:"type: varchar(225)"`
	Password  string          `json:"password" gorm:"type: varchar(225)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Role      string          `json:"-" gorm:"type: varchar(225)"`
}

//
type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
