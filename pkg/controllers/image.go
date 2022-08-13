package controllers

import (
	"fmt"
	"image-uploader/pkg/config"
	"image-uploader/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Image struct {
	URL string `json:"url"`
}

type ImageResponse struct {
	Image   Image  `json:"image"`
	Message string `json:"message"`
}

func UploadImage(context *gin.Context) {
	println("Upload")
	file, err := context.FormFile("file")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("[Error]: %s\n", err)
		return
	}

	fileName := fmt.Sprintf("%s.%s", utils.RandomString(12), file.Filename[strings.LastIndex(file.Filename, ".")+1:])

	fullPath := fmt.Sprintf("%s/%s", "public/images", fileName)

	err = context.SaveUploadedFile(file, fullPath)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Printf("[Error]: %s\n", err)
		return
	}

	response := ImageResponse{
		Image:   Image{URL: fmt.Sprintf("%s/public/images/%s", config.HOST, fileName)},
		Message: "uploaded",
	}
	context.JSON(http.StatusOK, response)

}
