package file

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func DeleteDir(path string) error {
	return os.RemoveAll(path)
}

func ListDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

type File struct {
	Path      string
	Content   []byte
	Size      int64
	Extension string
}

func ReadFile(path string) (*File, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &File{
		Path:      path,
		Content:   content,
		Size:      int64(len(content)),
		Extension: path[strings.LastIndex(path, ".")+1:],
	}, nil
}

// the function to check filename extension is valid from the provided list of extensions
func IsValidExtension(filename string, extensions []string) bool {
	ext := filepath.Ext(filename)
	for _, e := range extensions {
		if e == ext {
			return true
		}
	}
	return false
}

type FileResponse struct {
	Name      string
	Content   []byte
	Size      int64
	Extension string
}

func GetFileFromURL(url string) (*FileResponse, error) {
	// validate url
	if !strings.HasPrefix(url, "http") {
		return nil, errors.New("invalid url")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// check if content is a pdf or image
	if !IsValidExtension(filepath.Base(url), []string{"pdf", "jpg", "jpeg", "png", "gif"}) {
		return nil, errors.New("invalid file type")
	}

	return &FileResponse{
		Name:      filepath.Base(url),
		Content:   content,
		Size:      int64(len(content)),
		Extension: filepath.Ext(url),
	}, nil
}
