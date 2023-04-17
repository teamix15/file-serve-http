package main

import (
	"DocummentsServer/api/handler"
	"DocummentsServer/templates"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	templates.HandleTemplate()

	r.GET("/files", handler.GetFile)

	r.POST("/files", handler.UploadFile)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
