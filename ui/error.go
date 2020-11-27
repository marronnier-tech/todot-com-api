package ui

import (
	"github.com/gin-gonic/gin"
)

func errHundle(err error, status int, c *gin.Context) {

	if status == 500 {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

}
