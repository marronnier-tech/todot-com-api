package todo

import (
	"../../domain"
	"../../infra/table"
)

type allTodoArray struct {
	TodoObj domain.TodoObjInfo    `json:"TodoObj"`
	User    domain.UserSimpleInfo `json:"User"`
}

type userTodoArray struct {
	User    domain.UserSimpleInfo `json:"User"`
	TodoObj []domain.TodoObjInfo  `json:"TodoObj"`
}

type todayTodo struct {
	TodoLog       table.TodoAchievedLog `json:"TodoLogInfo"`
	TodayAchieved bool                  `json:"TodayAchieved"`
}

type OperationView struct {
	TodoID        int    `json:"TodoID"`
	IsDeleted     bool   `json:"IsDeleted"`
	Content       string `json:"Content"`
	CreatedAt     string `json:"CreatedAt"`
	LastAchieved  string `json:"LastAchieved"`
	Count         int64  `json:"Count"`
	TodayAchieved bool   `json:"TodayAchieved"`
}
