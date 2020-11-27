package main

import (
	"net/http"

	"./ui"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

func main() {

	r := gin.Default()
	// r.Static("../../../front/templates", "./../../../front/templates")
	r.StaticFS("/top", http.Dir("./../../../front/templates"))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("useradmin", store))

	todolist := r.Group("/todo")
	{
		todolist.GET("", ui.GetTodo)
		todolist.GET("/:name", ui.GetOneUserTodo)
	}

	goallist := r.Group("/goal")
	{
		goallist.GET("", ui.GetGoal)
		goallist.GET("/:name", ui.GetOneUserGoal)
	}

	my := r.Group("/mypage")
	{
		my.GET("", ui.MyTodo)
		my.POST("", ui.PostTodo)
		my.GET("/goal", ui.MyGoal)
		my.DELETE("/:id", ui.DeleteTodo)

		my.POST("/:id/today", ui.PutAchieveTodo)
		my.DELETE("/:id/today", ui.ClearAchieveTodo)
		my.PATCH("/:id/goal", ui.PatchGoal)

	}

	profile := r.Group("/profile")
	{
		profile.GET("", ui.GetMyProfile)
		profile.PATCH("", ui.PatchProfile)

		profile.GET("/:name", ui.GetOnesProfile)

	}

	r.POST("/register", ui.Register)

	r.POST("/login", ui.Login)
	r.DELETE("/logout", ui.Logout)
	r.GET("/adminflag", ui.AdminFlag)

	r.DELETE("/delete", ui.DeleteMembership)

	r.Run()

}
