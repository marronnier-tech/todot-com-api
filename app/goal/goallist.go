package goal

import (
	"../../domain"
	"../../infra"
	"../align"
	"../getid"
	"../timecalc"
)

func ToGetAllGoal(limit int, page int, order string) (out []allGoalArray, err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var rows []inGoal

	base := tx.Table("goal_lists").
		Select("goal_lists.ID, goal_lists.todo_id, goal_lists.count, goal_lists.goaled_at, todo_lists.Content, todo_lists.is_deleted, users.id as user_id, users.name, users.handle_name, users.img, users.goaled_count").
		Where("todo_lists.is_deleted = ?", false).
		Joins("join todo_lists on goal_lists.todo_id = todo_lists.id").
		Joins("join users on todo_lists.user_id = users.id").
		Limit(limit).
		Offset(limit * (page - 1))

	err = align.ListOrder(base, "goal_lists", true, order).
		Scan(&rows).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	var obj domain.OutGoalObjInfo
	var user domain.UserSimpleInfo

	for _, r := range rows {

		obj = domain.OutGoalObjInfo{
			TodoID:        r.ID,
			Content:       r.Content,
			GoaledAt:      timecalc.PickDate(r.GoaledAt),
			AchievedCount: r.Count,
		}

		if r.UserHN == nil {
			r.UserHN = &r.UserName
		}

		user = domain.UserSimpleInfo{
			UserID:      r.UserID,
			UserName:    r.UserName,
			UserHN:      r.UserHN,
			UserImg:     r.UserImg,
			GoaledCount: r.GoaledCount,
		}

		out = append(out, allGoalArray{
			GoalObj: obj,
			User:    user,
		})

	}

	err = tx.Commit().Error

	return

}

func ToGetOneGoal(name string, order string) (have bool, out userGoalArray, err error) {
	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	user, userID, err := getid.Fromname(tx, name)

	if err != nil {
		tx.Rollback()
		return
	}

	var rows []domain.GoalObjInfo

	base := tx.Table("todo_lists").
		Select("todo_lists.id as todo_id, todo_lists.content, goal_lists.goaled_at, goal_lists.count as achieved_count").
		Where("todo_lists.user_id = ? and todo_lists.is_deleted = ? and todo_lists.is_goaled = ?", userID, false, true).
		Joins("join goal_lists on todo_lists.id = goal_lists.todo_id")

	err = align.ListOrder(base, "goal_lists", true, order).
		Scan(&rows).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	var obj domain.OutGoalObjInfo
	var objArray []domain.OutGoalObjInfo

	for _, r := range rows {

		obj = domain.OutGoalObjInfo{
			TodoID:        r.TodoID,
			Content:       r.Content,
			GoaledAt:      timecalc.PickDate(r.GoaledAt),
			AchievedCount: r.AchievedCount,
		}

		objArray = append(objArray, obj)

	}

	out = userGoalArray{
		User:    user,
		GoalObj: objArray,
	}
	have = true

	err = tx.Commit().Error

	return

}
