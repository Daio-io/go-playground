package main

import "github.com/gin-gonic/gin"
import "go-test/handlers/home"

func main() {
	
	r:= gin.Default()
	
	r.GET("/", home.GetHome)
	
    r.GET("/user/:name", home.GetUser)
	
	r.Run(":3000")
	
}