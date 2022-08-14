package main

import (
	"fmt"
	"image-uploader/pkg/config"
	"image-uploader/pkg/controllers"
	"log"
	"net/http"
	"os"
	"time"

	corsgin "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	frontendHost := config.FT_HOST

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//config allow host
	router.SetTrustedProxies([]string{frontendHost})
	router.Use(corsgin.New(corsgin.Config{
		AllowOrigins:     []string{frontendHost},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/public/images", "./public/images")
	router.POST("/upload", controllers.UploadImage)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Vienvenido")
	})

	fmt.Println(frontendHost)

	router.Run(fmt.Sprintf(":%s", config.API_PORT))

}
