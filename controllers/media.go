package controllers

import (
	"io"
	"log"
	"mime/multipart"
	"mypress/database"
	"mypress/models"
	"mypress/utils"
	"os"
	"strings"
)

// CreateMediaPicture func
func CreateMediaPicture(file *multipart.FileHeader) uint {
	extension := strings.Split(file.Header["Content-Type"][0], "/")[1]
	fileName := strings.Join([]string{utils.UUID(), ".", extension}, "")
	content, err := file.Open()
	defer content.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	//content.Read(buf)

	filePath := strings.Join([]string{"./medias/", fileName}, "")
	outputFile, err := os.Create(filePath)
	defer outputFile.Close()

	for {
		// read a chunk
		n, err := content.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := outputFile.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	media := models.Media{Format: file.Header["Content-Type"][0], Name: fileName, Size: file.Size}
	database.Get().Create(&media)

	return media.ID
}
