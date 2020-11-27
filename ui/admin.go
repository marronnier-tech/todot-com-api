package ui

import (
	"errors"

	"../app/admin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionLogin(c *gin.Context) (id int, user string, err error) {

	session := sessions.Default(c)

	name := session.Get("name")
	password := session.Get("password")

	if name == nil || password == nil {
		err = errors.New("not found")
		c.JSON(404, err.Error)
		c.Abort()
		return
	}

	strname := name.(string)
	strpassword := password.(string)

	id, user, err = admin.Validation(strname, strpassword)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
	}
	return
}

func Login(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	session := sessions.Default(c)

	_, _, err := admin.Validation(name, password)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session.Set("name", name)
	session.Set("password", password)
	session.Save()

	// c.Redirect(302, "/mypage")
	c.JSON(201, nil)
	return

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(204, gin.H{"message": "Logout done successful"})
}

func Register(c *gin.Context) {
	session := sessions.Default(c)

	name := c.PostForm("name")
	password := c.PostForm("password")

	if err := admin.SignUp(name, password); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session.Set("name", name)
	session.Set("password", password)
	session.Save()

	c.Redirect(302, "/mypage")

}

func DeleteMembership(c *gin.Context) {
	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = admin.ToDeleteMember(userid)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"DeleteProcess": true,
	})

	c.Redirect(302, "/todo")

}

func AdminFlag(c *gin.Context) {
	session := sessions.Default(c)

	name := session.Get("name")
	password := session.Get("password")

	flag := true

	if name == nil || password == nil {
		flag = false
	}

	c.JSON(200, gin.H{
		"LoginFlag": flag,
	})

}
