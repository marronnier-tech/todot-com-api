package goal

import (
	"github.com/tocchy-tocchy/todot-com-api/domain"
)

type allGoalArray struct {
	GoalObj domain.OutGoalObjInfo `json:"GoalObj"`
	User    domain.UserSimpleInfo `json:"User"`
}

type userGoalArray struct {
	User    domain.UserSimpleInfo   `json:"User"`
	GoalObj []domain.OutGoalObjInfo `json:"GoalObj"`
}
