package goal

import (
	"../../infra/table"
)

type inGoal struct {
	table.GoalList
	Content     string  `gorm:"column:content"`
	UserID      int     `gorm:"column:user_id"`
	UserName    string  `gorm:"column:name"`
	UserHN      *string `gorm:"column:name"`
	UserImg     *string `gorm:"column:img"`
	GoaledCount int64   `gorm:"column:goaled_count"`
}
