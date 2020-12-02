package table

import "time"

type FollowList struct {
	ID         int       `gorm:"id; autoIncrement"`
	UserID     int       `gorm:"user_id"`
	FollowedID int       `gorm:"follwed_id"`
	IsDeleted  bool      `gorm:"is_deleted;default:false"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
}
