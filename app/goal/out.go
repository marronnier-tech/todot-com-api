package goal

import (
	"../../domain"
)

type allGoalArray struct {
	GoalObj domain.OutGoalObjInfo `json:"GoalObj"`
	User    domain.UserSimpleInfo `json:"User"`
}

type userGoalArray struct {
	User    domain.UserSimpleInfo   `json:"User"`
	GoalObj []domain.OutGoalObjInfo `json:"GoalObj"`
}
