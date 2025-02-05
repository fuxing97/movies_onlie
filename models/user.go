package models

type User struct {
	ID           int    `gorm:"type:bigint(0);not null" json:"id" binding:"required"`
	UserName     string `gorm:"type:varchar(255);not null" json:"user_name" binding:"required"`
	UserPassword string `gorm:"type:char(32);not null" json:"user_password" binding:"required"`
	UserType     string `gorm:"type:tinyint(1);not null" json:"user_type" binding:"required"`
	Gender       string `gorm:"type:tinyint(1);not null" json:"gender" binding:"required"`
	Email        string `gorm:"type:varchar(255);not null" json:"email" binding:"required"`
	IsAdmin      string `gorm:"type:tinyint(1);not null" json:"is_admin" binding:"required"`
}

func (User) TableName() string {
	return "user"
}

type LoginUser struct {
	Name     string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	PassWord string `gorm:"type:char(32);not null" json:"passwd" binding:"required"`
}

type RegisterUser struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"password"`
	UserType     string `json:"user_type"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
}
