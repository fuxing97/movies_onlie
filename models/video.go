package models

type MovieInfo struct{
	ID           int    `gorm:"type:bigint(0);not null" json:"id" binding:"required"`
	MovieID     string `gorm:"type:varchar(32);not null" json:"movie_id" binding:"required"`
	MovieName string `gorm:"type:varchar(255);not null" json:"movie_name" binding:"required"`
	Category     string `gorm:"type:varchar(255);not null" json:"category" binding:"required"`
	ImgCover       string `gorm:"type:text(0);not null" json:"img_cover" binding:"required"`
	Actors        string `gorm:"type:text(0);not null" json:"actors" binding:"required"`
	Directors      string `gorm:"type:text(0);not null" json:"directors" binding:"required"`
	Year string `gorm:"type:varchar(255);not null" json:"year" binding:"required"`
	Language string `gorm:"type:varchar(64);not null" json:"language" binding:"required"`
	Duration string `gorm:"type:varchar(32);not null" json:"duration" binding:"required"`
	Description string `gorm:"type:text(0);not null" json:"description" binding:"required"`
	Country string `gorm:"type:varchar(100);not null" json:"country" binding:"required"`
	Score string `gorm:"type:char(115);not null" json:"score" binding:"required"`
	UpdateInfo string `gorm:"type:varchar(255);not null" json:"update_info" binding:"required"`
	Status 	int `gorm:"type:tinyint(1);not null" json:"status" binding:"required"`
	MovieType string `gorm:"type:varchar(30);not null" json:"movie_type" binding:"required"`
	IsFinish int `gorm:"type:tinyint(1);not null" json:"is_finish" binding:"required"`
}

func (MovieInfo) TableName() string {
	return "movie_info"
}