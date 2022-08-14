package main

import (
	"bytes"
	"fmt"
	"image-uploader/pkg/config"
	"image-uploader/pkg/controllers"
	"net/http"
	"time"

	corsgin "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from uploader images!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}
}

func main() {

	router := gin.Default()

	frontendHost := config.FT_HOST

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
	fmt.Println(frontendHost)
	router.Run(fmt.Sprintf(":%s", config.API_PORT))
}
