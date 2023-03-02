package models

import "time"

// product modal struct
type Product struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name        string               `json:"name" gorm:"type: varchar(225)"`
	Price       int                  `json:"price" gorm:"type:int"`
	Description string               `json:"description" gorm:"type: varchar(225)"`
	Stock       int                  `json:"stock"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	UserID      int                  `json:"user_id" form:"user_id"`
	User        UsersProfileResponse `json:"user"`
	Image       string               `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type ProductResponse struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Desc   string               `json:"desc"`
	Price  int                  `json:"price"`
	Image  string               `json:"image"`
	Stock  int                  `json:"stock"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}
