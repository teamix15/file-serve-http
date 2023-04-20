package handler

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(c *gin.Context) {
	filename := c.Query("filename")
	filePath := filepath.Join("data", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
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
		path = "data"
	}

	if err := c.SaveUploadedFile(file, path+"/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File has been successfully uploaded",
	})
}

func GetFileNames(c *gin.Context) {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read directory")
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{
		"files": fileNames,
	})
}

func DeleteAllFiles(c *gin.Context) {
	files, err := os.ReadDir("data")
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read directory")
		return
	}

	for _, file := range files {
		err = os.Remove(filepath.Join("data", file.Name()))
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not remove file")
			return
		}
	}

	c.String(http.StatusOK, "All files removed")
}

func DeleteFile(c *gin.Context) {
	filename := c.PostForm("filename")

	if err := os.Remove(filepath.Join("data", filename)); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "File deleted successfully"})
}

func DownloadAllFiles(c *gin.Context) {
	directory := "data"

	zipFileName := "files.zip"

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		zipFileInArchive, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFileInArchive, file)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.String(500, err.Error())
		return
	}

	err = zipWriter.Close()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "application/zip")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", zipFileName))
	c.File(zipFileName)
}
