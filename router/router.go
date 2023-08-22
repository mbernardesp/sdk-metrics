package router

import "github.com/gin-gonic/gin"

func Init() {
	//Initialize Router
	router := gin.Default()

	initializeRoutes(router)

	//Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
