package main

import (
	"bytes"
	"fmt"
	"image-uploader/pkg/config"
	"image-uploader/pkg/controllers"
	"log"
	"net/http"
	"os"
	"strconv"
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

	frontendHost := config.FT_HOST

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	tStr := os.Getenv("REPEAT")
	repeat, err := strconv.Atoi(tStr)

	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
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
	fmt.Println(frontendHost)

	router.GET("/repeat", repeatHandler(repeat))
	router.Run(fmt.Sprintf(":%s", config.API_PORT))

}
