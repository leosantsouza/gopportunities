package router

import "github.com/gin-gonic/gin"

func Initialize() {

//Initializer Router
	router := gin.Default()
	
// Initialize Routes
	initializeRoutes(router)

// Run the Server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
