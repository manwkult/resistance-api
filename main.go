package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "resistance/docs"
)

// @title           Swagger Resistance API
// @version         1.0
// @description     API para interceptar los mensajes encriptador recibidos por los satelites
// @termsOfService  http://swagger.io/terms/

// @contact.name   	API Resistance
// @contact.url    	http://www.swagger.io/support
// @contact.email  	manuel_ldsc@hotmail.com

// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @host      		http://143.244.202.236
// @BasePath  		/api/v1
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	var url = ginSwagger.URL("http://localhost:8000/swagger/doc.json")

	if os.Getenv("ENV") == "production" {
		url = ginSwagger.URL("http://143.244.202.236/swagger/doc.json")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	ping := router.Group("/api/ping")
	{
		ping.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "Ok"})
		})
	}

	v1 := router.Group("/api/v1/")
	{
		v1.POST("/topsecret", topSecret)
		v1.GET("/topsecret_split", topSecretSplitGET)
		v1.POST("/topsecret_split", topSecretSplitPOST)
	}

	router.Run(":8000")
}
