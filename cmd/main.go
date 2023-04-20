package main

import (
	"DocummentsServer/api/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/files/download/", handler.DownloadFile)

	router.POST("/files/upload", handler.UploadFile)

	router.GET("/files/names", handler.GetFileNames)

	router.POST("/files/delete/all", handler.DeleteAllFiles)

	router.POST("/files/delete/", handler.DeleteFile)

	router.GET("/files/download/all", handler.DownloadAllFiles)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
