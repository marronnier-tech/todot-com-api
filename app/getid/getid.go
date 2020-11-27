package getid

import (
	"../../domain"
	"../../infra/table"
	"gorm.io/gorm"
)

func Fromname(tx *gorm.DB, name string) (user domain.UserSimpleInfo, userID int, err error) {

	var u table.User

	err = tx.Table("users").
		Where("name = ?", name).
		Scan(&u).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	userID = u.ID

	if u.HN == nil {
		u.HN = &u.Name
	}

	user = domain.UserSimpleInfo{
		UserID:      u.ID,
		UserName:    u.Name,
		UserHN:      u.HN,
		UserImg:     u.Img,
		GoaledCount: u.GoaledCount,
	}

	return
}
