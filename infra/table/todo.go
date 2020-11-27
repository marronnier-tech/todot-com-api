package table

import (
	"time"

	"github.com/lib/pq"
)

type TodoList struct {
	ID           int         `gorm:"column:id;autoIncrement"`
	UserID       int         `gorm:"column:user_id"`
	Content      string      `gorm:"column:content"`
	CreatedAt    time.Time   `gorm:"column:created_at"`
	LastAchieved pq.NullTime `gorm:"column:last_achieved"`
	Count        int64       `gorm:"column:count;default:0"`
	IsDeleted    bool        `gorm:"column:is_deleted;default:false"`
	IsGoaled     bool        `gorm:"column:is_goaled;default:false"`
}

// 達成ログ

type TodoAchievedLog struct {
	ID           int         `gorm:"column:id; autoIncrement"`
	TodoID       int         `gorm:"column:todo_id"`
	AchievedDate pq.NullTime `gorm:"column:achieved_date"`
	IsDeleted    bool        `gorm:"column:is_deleted;default:false"`
}

// ゴールログ

type GoalList struct {
	ID       int       `gorm:"column:id; autoIncrement"`
	TodoID   int       `gorm:"column:todo_id"`
	Count    int64     `gorm:"column:count"`
	GoaledAt time.Time `gorm:"column:goaled_at"`
}
