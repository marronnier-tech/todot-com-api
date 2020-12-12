package main

import (
	"net/http"

	"github.com/tocchy-tocchy/todot-com-api/ctrler"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

func main() {

	r := gin.Default()
	// r.Static("../../../front/templates", "./../../../front/templates")
	r.StaticFS("/top", http.Dir("./../../../front/templates"))

	r.Use(cors.New(ctrler.CorsConfig()))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("useradmin", store))

	todolist := r.Group("/todo")
	{
		todolist.GET("", ctrler.GetTodo)
		todolist.GET("/:name", ctrler.GetOneUserTodo)
	}

	goallist := r.Group("/goal")
	{
		goallist.GET("", ctrler.GetGoal)
		goallist.GET("/:name", ctrler.GetOneUserGoal)
	}

	my := r.Group("/mypage")
	{
		my.GET("", ctrler.MyTodo)
		my.POST("", ctrler.PostTodo)
		my.GET("/goal", ctrler.MyGoal)
		my.DELETE("/:id", ctrler.DeleteTodo)

		my.POST("/:id/today", ctrler.PutAchieveTodo)
		my.DELETE("/:id/today", ctrler.ClearAchieveTodo)
		my.PATCH("/:id/goal", ctrler.PatchGoal)

	}

	profile := r.Group("/profile")
	{
		profile.GET("", ctrler.GetMyProfile)
		profile.PATCH("", ctrler.PatchProfile)

		profile.GET("/:name", ctrler.GetOnesProfile)

	}

	r.POST("/register", ctrler.Register)

	r.POST("/login", ctrler.Login)
	r.DELETE("/logout", ctrler.Logout)
	r.GET("/adminflag", ctrler.AdminFlag)

	r.DELETE("/delete", ctrler.DeleteMembership)

	r.Run()

}
