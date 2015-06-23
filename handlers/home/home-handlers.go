package home

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	
	name := c.Params.ByName("name")
		
	c.JSON(200, gin.H{"user": name, "value": "im a user"})
		
}

func GetHome(c *gin.Context) {

	c.String(200, "pong")

}