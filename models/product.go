package models

import "time"

type Product struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: varchar(255)"`
	Price       int                  `json:"price" gorm:"type: int"`
	Stock       int                  `json:"stock" gorm:"type: int"`
	Photo       string               `json:"photo" gorm:"type: varchar(255)"`
	UserID      int                  `json:"user_id" form:"user_id"`
	User        UsersProfileResponse `json:"user"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type ProductResponse struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name" form:"name" validate:"required"`
	Price       int                  `json:"price" form:"price" validate:"required"`
	Description string               `json:"description" form:"description" validate:"required"`
	Stock       int                  `json:"stock" form:"stock" validate:"required"`
	Photo       string               `json:"photo" form:"photo" validate:"required"`
	UserID      int                  `json:"-"`
	User        UsersProfileResponse `json:"user"`
}

type ProductUserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Photo       string `json:"photo"`
	UserID      int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
