package ui

import (
	"fmt"

	"../app/admin"
	"../app/profile"
	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) {
	_, name, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, fmt.Sprintf("/profile/%s", name))

}

func PatchProfile(c *gin.Context) {
	userid, _, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	HN := c.PostForm("handle_name")
	Img := c.PostForm("img")
	FinalGoal := c.PostForm("final_goal")
	Profile := c.PostForm("profile")
	Twitter := c.PostForm("twitter")
	Instagram := c.PostForm("instagram")
	Facebook := c.PostForm("facebook")
	Github := c.PostForm("github")
	URL := c.PostForm("url")

	err = profile.ToPatch(userid, HN, Img, FinalGoal,
		Profile, Twitter, Instagram, Facebook, Github, URL)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, nil)

	return
}

func GetOnesProfile(c *gin.Context) {
	_, user, err := SessionLogin(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")
	res, err := profile.ToGetOneProfile(name)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"UserInfo": res,
		"owner":    admin.JudgeOwner(user, name),
	})

}
