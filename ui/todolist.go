package ui

import (
	"fmt"
	stc "strconv"

	"../app/admin"
	"../app/todo"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {

	page, _ := stc.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := stc.Atoi(c.DefaultQuery("limit", "100"))
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetAll(limit, page, order)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"TodoArray": res,
		"limit":     limit,
		"page":      page,
		"order":     order,
	})

	return

}

func GetOneUserTodo(c *gin.Context) {

	_, user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")
	order := c.DefaultQuery("order", "last_achieved")

	res, err := todo.ToGetOneUser(name, order)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Todo":  res,
		"order": order,
		"owner": admin.JudgeOwner(user, name),
	})

}

func MyTodo(c *gin.Context) {
	_, name, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	order := c.DefaultQuery("order", "last_achieved")
	c.Redirect(302, fmt.Sprintf("/todo/%s?order=%s", name, order))

}
