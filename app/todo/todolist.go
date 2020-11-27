package todo

import (
	"../align"
	"../getid"
	"../timecalc"

	"../../domain"
	"../../infra"
	"../../infra/table"
)

func ToGetAll(limit int, page int, order string) (out []allTodoArray, err error) {

	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var rows []inGetAll

	base := tx.Table("todo_lists").
		Select("todo_lists.id, todo_lists.Content, todo_lists.user_id, todo_lists.created_at, todo_lists.last_achieved, todo_lists.count, todo_lists.is_deleted, todo_lists.is_goaled, users.name, users.handle_name, users.img, users.goaled_count").
		Where("todo_lists.is_deleted = ? and todo_lists.is_goaled = ? and users.is_deleted = ?", false, false, false).
		Joins("left join users on users.ID = todo_lists.user_id").
		Limit(limit).
		Offset(limit * (page - 1))

	err = align.ListOrder(base, "todo_lists", true, order).
		Scan(&rows).
		Error

	if err != nil {
		return
	}

	var obj domain.TodoObjInfo
	var user domain.UserSimpleInfo

	for _, r := range rows {

		obj = domain.TodoObjInfo{
			TodoID:      r.ID,
			IsDeleted:   r.IsDeleted,
			Content:     r.Content,
			CreatedAt:   timecalc.PickDate(r.CreatedAt),
			Count:       r.Count,
			AchieveInfo: timecalc.DifftoNow(r.LastAchieved),
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

		out = append(out, allTodoArray{
			TodoObj: obj,
			User:    user,
		})

	}

	err = tx.Commit().Error

	return

}

func ToGetOneUser(name string, order string) (out userTodoArray, err error) {
	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	user, userID, err := getid.Fromname(tx, name)

	if err != nil {
		tx.Rollback()
		return
	}

	var rows []table.TodoList

	base := tx.Table("todo_lists").
		Select("id, user_id, content, created_at, last_achieved, count, is_deleted, is_goaled").
		Where("user_id = ? and is_deleted = ? and is_goaled = ?", userID, false, false)

	err = align.ListOrder(base, "todo_lists", false, order).
		Scan(&rows).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	var obj domain.TodoObjInfo
	var objArray []domain.TodoObjInfo

	for _, r := range rows {

		obj = domain.TodoObjInfo{
			TodoID:      r.ID,
			IsDeleted:   r.IsDeleted,
			Content:     r.Content,
			CreatedAt:   timecalc.PickDate(r.CreatedAt),
			Count:       r.Count,
			AchieveInfo: timecalc.DifftoNow(r.LastAchieved),
		}

		objArray = append(objArray, obj)

	}

	out = userTodoArray{
		User:    user,
		TodoObj: objArray,
	}

	err = tx.Commit().Error

	return

}
