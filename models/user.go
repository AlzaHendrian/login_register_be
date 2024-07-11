package models

import "time"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"name" gorm:"type: varchar(50)"`
	LastName    string `json:"last_name" gorm:"type: varchar(50)"`
	DateOfBirth string `json:"date_of_birth" gorm:"type: varchar(50)"`
	Gender      string `json:"gender" gorm:"type: varchar(50)"`
	Email       string `json:"email" gorm:"type: varchar(255)"`
	Password    string `json:"password" gorm:"type: varchar(255)"`
	Role        string `json:"role" gorm:"type: varchar(50)"`
	IsAdmin     bool   `json:"isAdmin" gorm:"default:false"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
