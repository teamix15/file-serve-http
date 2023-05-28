package main

import (
	"DocummentsServer/article/delivery"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/files/download/", delivery.DownloadFile)

	router.POST("/files/upload", delivery.UploadFile)

	router.GET("/files/names", delivery.GetFileNames)

	router.POST("/files/delete/all", delivery.DeleteAllFiles)

	router.POST("/files/delete/", delivery.DeleteFile)

	//router.GET("/files/download/all", delivery.DownloadAllFiles)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
