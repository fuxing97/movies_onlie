package models

type VideoAction struct {
	ID           int `gorm:"type:bigint(0);not null" json:"id" binding:"required"`
	VID          int `gorm:"type:bigint(0);not null" json:"v_id" binding:"required"`
	LikeCount    int `gorm:"type:bigint(0);not null" json:"like_count" binding:"required"`
	DislikeCount int `gorm:"type:bigint(0);not null" json:"dislike_count" binding:"required"`
	PlayCount    int `gorm:"type:bigint(0);not null" json:"play_count" binding:"required"`
}

func (VideoAction) TableName() string {
	return "video_action"
}
