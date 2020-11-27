package todo

import (
	"../../infra/table"
)

type inGetAll struct {
	table.TodoList
	UserName    string  `gorm:"column:name"`
	UserHN      *string `gorm:"column:handle_name"`
	UserImg     *string `gorm:"column:img"`
	GoaledCount int64   `gorm:"column:goaled_count"`
}

type sameCheck struct {
	IsDeleted bool   `gorm:"column:is_deleted"`
	Content   string `gorm:"column:content"`
}
