package infra

import (
	"fmt"

	"./def"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (tx *gorm.DB, err error) {

	env, err := def.GetEnv()

	if err != nil {
		return
	}

	connect := fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s",
		env.User, env.Pass, def.Protocol, def.Database, def.Charset, def.ParseTime, def.Loc)
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
