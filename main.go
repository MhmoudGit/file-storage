package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const Path = "./storage/"

type File struct {
	Name string
	Size int
	Path string
	Type string
}

type Storage struct {
	Name  string
	Path  string
	Files []File
}

type User struct {
	Username    string
	StorageSize int
	UserSpace   Storage
}

func CreateUser(username string) User {
	storageSizeInKb := 1024 * 1024 * 100
	spaceName := username + "-space"
	err := os.Mkdir(Path+spaceName, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	return User{
		Username:    username,
		StorageSize: storageSizeInKb,
		UserSpace: Storage{
			Name: spaceName,
			Path: Path + spaceName + "/",
		},
	}
}

func (u *User) UploadFile(file *os.File) {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	u.StorageSize -= int(fileInfo.Size())

	if u.StorageSize <= 0 {
		log.Fatal("Looks like you are out of storage")
		return
	}
	filePath := filepath.Join(u.UserSpace.Path, file.Name())
	filePtr, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
	// Copy the uploaded file data to the new file
	_, err = io.Copy(filePtr, file)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the file type (MIME type)
	fileHeader := make([]byte, 512) // Read the first 512 bytes to detect the content type
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Read(fileHeader)
	if err != nil {
		log.Fatal(err)
	}
	fileType := http.DetectContentType(fileHeader)

	u.UserSpace.Files = append(u.UserSpace.Files, File{
		Name: fileInfo.Name(),
		Size: int(fileInfo.Size()),
		Path: filePath,
		Type: fileType,
	})
}

func main() {
	mahmoud := CreateUser("mo")

	inputFile, err := os.Open("./image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	mahmoud.UploadFile(inputFile)

	fmt.Println(mahmoud)
}
