package ui

import (
	"fmt"
	stc "strconv"

	"../app/admin"
	"../app/goal"
	"github.com/gin-gonic/gin"
)

func GetGoal(c *gin.Context) {
	page, _ := stc.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := stc.Atoi(c.DefaultQuery("limit", "100"))
	order := c.DefaultQuery("order", "goaled_at")

	res, err := goal.ToGetAllGoal(limit, page, order)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"GoalArray": res,
		"limit":     limit,
		"page":      page,
		"order":     order,
	})

}

func GetOneUserGoal(c *gin.Context) {

	_, user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")
	order := c.DefaultQuery("order", "goaled_at")

	have, res, err := goal.ToGetOneGoal(name, order)

	if have == false {
		c.JSON(404, gin.H{"message": "このユーザーにはゴールしたTODOがありません"})
		return

	}

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return

	}

	c.JSON(200, gin.H{
		"Goal":  res,
		"order": order,
		"owner": admin.JudgeOwner(user, name),
	})

}

func MyGoal(c *gin.Context) {
	_, name, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	order := c.DefaultQuery("order", "goaled_at")

	c.Redirect(302, fmt.Sprintf("/goal/%s?order=%s", name, order))

}
