package delivery

import (
	"DocummentsServer/article/repository"
	"DocummentsServer/view/templates"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(c *gin.Context) {
	filename := c.Query("filename")
	filePath := filepath.Join("data", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		c.JSON(http.StatusBadRequest, "Invalid file format. Only JPG, PNG, and GIF are allowed.")
		return
	}

	if err = repository.AddNewFile(file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(file, filepath.Join("data", file.Filename)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "File has been successfully uploaded")
}

func GetFileNames(c *gin.Context) {
	files, err := repository.GetFileNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	html := templates.GenerateHTMLFilesList(files)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}

func DeleteAllFiles(c *gin.Context) {
	files, err := os.ReadDir("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, file := range files {
		if err = os.Remove(filepath.Join("data", file.Name())); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	if err = repository.DeleteAllFiles(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "All files have been successfully deleted")
}

func DeleteFile(c *gin.Context) {
	filename := c.PostForm("filename")

	if err := os.Remove(filepath.Join("data", filename)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := repository.DeleteFile(filename); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "File deleted successfully")
}
