package infra

import (
	"fmt"

	"github.com/tocchy-tocchy/todot-com-api/infra/def"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (tx *gorm.DB, err error) {

	env, err := def.GetEnv()

	if err != nil {
		return
	}

	var connect string

	if env.JawsDB != "" {
		checker := 0
		var sqlUser string
		var sqlPass string
		var sqlHost string
		var sqlDatabase string

		for i, jw := range env.JawsDB {
			if i < 7 {
				continue
			}
			if checker == 0 {
				if jw == ':' {
					checker++
					continue
				}
				sqlUser += string(jw)
			} else if checker == 1 {
				if jw == '@' {
					checker++
					continue
				}
				sqlPass += string(jw)
			} else if checker == 2 {
				if jw == '/' {
					checker++
					continue
				}
				sqlHost += string(jw)
			} else if checker == 3 {
				if jw == '?' {
					break
				}
				sqlDatabase += string(jw)
			}

		}

		connect = fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?parseTime=%s",
			string(sqlUser), string(sqlPass), string(sqlHost), string(sqlDatabase), def.ParseTime)

	} else {
		connect = fmt.Sprintf(
			"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
			env.User, env.Pass, def.Protocol, def.Database, def.Charset, def.ParseTime, def.Loc)
	}
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	sqlDB, err := db.DB()

	if err != nil {
		return
	}

	tx = db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
		sqlDB.Close()
	}()

	err = tx.Error

	if err != nil {
		return
	}

	return
}
