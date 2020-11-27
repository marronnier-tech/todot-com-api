package admin

import (
	"fmt"
	"time"

	"../../domain"
	"../../infra"
	"../../infra/table"
	"golang.org/x/crypto/bcrypt"
)

func Validation(name string, password string) (id int, user string, err error) {
	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var userauth domain.LoginInfo

	tx.Table("users").Select("id, name, password").Where("name = ?", name).Find(&userauth)
	selectpass := userauth.Password

	err = bcrypt.CompareHashAndPassword(selectpass, []byte(password))

	if err != nil {
		tx.Rollback()
		return
	}

	id = userauth.UserID
	user = userauth.UserName

	tx.Commit()
	return
}

func SignUp(user string, password string) (err error) {

	tx, err := infra.DBConnect()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(hash)

	if err != nil {
		tx.Rollback()
		return err
	}

	newuser := table.User{
		Name:        user,
		Password:    hash,
		MailAddress: nil,
		HN:          nil,
		Img:         nil,
		FinalGoal:   nil,
		Profile:     nil,
		Twitter:     nil,
		Instagram:   nil,
		Facebook:    nil,
		Github:      nil,
		URL:         nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	err = tx.Create(&newuser).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func ToDeleteMember(userid int) (err error) {
	tx, err := infra.DBConnect()

	if err != nil {
		return
	}

	var user table.User

	err = tx.Table("users").
		Where("id = ?", userid).
		First(&user).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Model(&user).Update("is_deleted", true)

	var todo table.TodoList

	err = tx.Table("todo_lists").
		Where("user_id = ?", userid).
		Find(&todo).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	todo.IsDeleted = true
	tx.Save(&todo)

	return tx.Commit().Error

}

func JudgeOwner(loginuser string, name string) bool {

	if loginuser == name {
		return true
	}

	return false

}
