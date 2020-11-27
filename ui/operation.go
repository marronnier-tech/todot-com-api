package ui

import (
	stc "strconv"

	"../app/todo"

	"github.com/gin-gonic/gin"
)

func PostTodo(c *gin.Context) {
	id, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	content := c.PostForm("content")

	if content == "" {
		c.JSON(500, gin.H{"error": "ToDoに設定したい内容を入力してください"})
		return
	}

	res, err := todo.ToPost(id, content)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"TodoObj": res})

	return

}

func DeleteTodo(c *gin.Context) {
	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	todoid, _ := stc.Atoi(c.Param("id"))

	err = todo.ToDelete(todoid, userid)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, nil)

	return
}

func PutAchieveTodo(c *gin.Context) {

	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	todoid, _ := stc.Atoi(c.Param("id"))

	res, err := todo.ToPutAchieve(todoid, userid)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

func ClearAchieveTodo(c *gin.Context) {

	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	todoid, _ := stc.Atoi(c.Param("id"))

	res, err := todo.ToClearAchieve(todoid, userid)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func PatchGoal(c *gin.Context) {
	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	todoid, _ := stc.Atoi(c.Param("id"))

	err = todo.ToPatchGoal(todoid, userid)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, nil)

}
