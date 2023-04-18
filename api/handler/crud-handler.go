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

const (
	storagePath = "/home/dad/Documents/Golang/file-serve-http/data"
)

func GetFile(c *gin.Context) {
	filename := c.Query("filename")
	filePath := "data/" + filename

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}

func GetAllFiles(c *gin.Context) {
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

func GetFilenames(c *gin.Context) {
	dir := storagePath
	files, err := ioutil.ReadDir(dir)
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

func DeleteAllFiles(c *gin.Context) {
	dir := storagePath
	files, err := os.ReadDir(dir)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read directory")
		return
	}

	for _, file := range files {
		err = os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not remove file")
			return
		}
	}

	c.String(http.StatusOK, "All files removed")
}

func DeleteFile(c *gin.Context) {
	filename := c.Query("filename")

	err := os.Remove(filename)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Не удалось удалить файл: %s", filename)})
		return
	}

	c.JSON(200, gin.H{"message": "File delete successfully"})
}
