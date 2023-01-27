package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"segi3.com/api/pkg/books"
	"segi3.com/api/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"https://example.com"}
	corsConfig.AllowCredentials = false

	router.Use(cors.New(corsConfig))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.Run(port)

}
