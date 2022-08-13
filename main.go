package main

import (
	"fmt"
	"image-uploader/pkg/config"
	"image-uploader/pkg/controllers"
	"time"

	corsgin "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	frontendHost := config.FT_HOST

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.SetTrustedProxies([]string{frontendHost})
	router.Use(corsgin.New(corsgin.Config{
		AllowOrigins:     []string{frontendHost},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/public/images", "./public/images")
	router.POST("/upload", controllers.UploadImage)
	fmt.Println(frontendHost)
	router.Run(fmt.Sprintf(":%s", config.API_PORT))
}
