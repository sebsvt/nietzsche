package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	uploadDir string
	baseURL   string
}

func NewLocalStorage(uploadDir, baseURL string) (FileStorage, error) {
	// check if uploadDir exists and create it if it doesn't
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &LocalStorage{
		uploadDir: uploadDir,
		baseURL:   baseURL,
	}, nil
}

// Save implements FileStorage.
func (l *LocalStorage) Save(documentID string, content io.Reader) (string, error) {
	// Create the file
	filePath := filepath.Join(l.uploadDir, documentID)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// copy the content to the file
	_, err = io.Copy(file, content)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// Delete implements FileStorage.
func (l *LocalStorage) Delete(documentID string) error {
	filePath := filepath.Join(l.uploadDir, documentID)
	return os.Remove(filePath)
}

// Get implements FileStorage.
func (l *LocalStorage) Get(documentID string) (io.ReadCloser, error) {
	filePath := filepath.Join(l.uploadDir, documentID)
	return os.Open(filePath)
}

// URL implements FileStorage.
func (l *LocalStorage) URL(documentID string) (string, error) {
	return fmt.Sprintf("%s/%s", l.baseURL, documentID), nil
}

// Close implements FileStorage.
func (l *LocalStorage) Close() error {
	return nil
}
