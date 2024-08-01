package main

import "github.com/gin-gonic/gin"

func main()  {
	
	router := gin.Default()

  router.GET("/ping", handlers...:func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run(addr...: ":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}