package models

type VideoInfo struct{
	ID           int    `gorm:"type:bigint(0);not null" json:"id" binding:"required"`
	MovieID     string `gorm:"type:varchar(32);not null" json:"movie_id" binding:"required"`
	VideoName string `gorm:"type:varchar(255);not null" json:"video_name" binding:"required"`
	SourceType     string `gorm:"type:tinyint(1);not null" json:"source_type" binding:"required"`
	SourceURL       string `gorm:"type:text(0);not null" json:"source_url" binding:"required"`
	
}

type VideoList struct {
	ID           int    `gorm:"type:bigint(0);not null" json:"id" binding:"required"`
	MovieID     string `gorm:"type:varchar(32);not null" json:"movie_id" binding:"required"`
	VideoName string `gorm:"type:varchar(255);not null" json:"video_name" binding:"required"`
}

func (VideoList) TableName() string {
	return "video_info"
}

func (VideoInfo) TableName() string {
	return "video_info"
}