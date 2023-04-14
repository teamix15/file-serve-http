package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

const (
	storagePath = "/home/dad/Documents/Golang/file-serve-http/data"
)

func GetFile(c *gin.Context) {
	filename := c.Query("filename")

	destination := c.Query("destination")

	source, err := os.Open(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer source.Close()

	newFile, err := os.Create(destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s copied successfully to %s", filename, destination),
	})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	path := c.PostForm("path")
	if path == "" {
		path = storagePath
	}

	if err := c.SaveUploadedFile(file, path+"/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", file.Filename),
	})
}
