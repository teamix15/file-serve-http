package main

import (
	"DocummentsServer/api/handler"
	"DocummentsServer/templates"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	templates.HandleTemplate()

	router.GET("/files/download/", handler.GetFile)

	router.GET("/files/names", handler.GetFilenames)

	router.POST("/files/upload", handler.UploadFile)

	router.POST("/files/delete/all", handler.DeleteAllFiles)

	router.POST("/files/delete/", handler.DeleteFile)

	router.GET("/files/download/all", handler.GetAllFiles)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
