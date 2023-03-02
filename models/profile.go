package models

import "time"

type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Phone     int                  `json:"phone" gorm:"type:int"`
	Gender    string               `json:"gender" gorm:"type:varchar(225)"`
	Address   string               `json:"address" gorm:"type:varchar(225)"`
	Photo     string               `json:"photo" gorm:"type:varchar(225)"`
	UserId    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"users"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type ProfileResponse struct {
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	Photo   string `json:"-"`
	UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
