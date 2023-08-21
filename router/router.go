package router

import "github.com/gin-gonic/gin"

func Init() {
	//Inicializa o router utilizando as configurações padrões do gin
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Executando a api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
