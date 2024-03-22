//storage: This folder contains modules or utilities related to data storage.

//fileStorage.go: Contains functions to interact with file storage systems, such as uploading files to disk or cloud storage.

package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// FileStorage provides functions for file storage operations.
type FileStorage struct {
	UploadDirectory string
}

// NewFileStorage creates a new FileStorage instance.
func NewFileStorage(uploadDirectory string) *FileStorage {
	return &FileStorage{
		UploadDirectory: uploadDirectory,
	}
}

// UploadFile uploads a file to the specified directory.
func (fs *FileStorage) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	filePath := filepath.Join(fs.UploadDirectory, fileHeader.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return filePath, nil
}

// DownloadFile downloads a file from the specified path.
func (fs *FileStorage) DownloadFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
