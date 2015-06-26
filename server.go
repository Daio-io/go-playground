package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "go-test/handlers"
)

// support for working on Google App Engine
 func init() {

  r:= gin.Default()

  r.GET("/", handlers.GetHome)
	
  r.GET("/user/:name", handlers.GetUser)

  http.Handle("/", r)
	
 }

// Local run
//  func main() {
	
// 	r:= gin.Default()

// 	r.GET("/", home.GetHome)
	
// 	r.GET("/user/:name", home.GetUser)
	
// 	r.Run(":3000")
	
// }