package nietzsche

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/sebsvt/nietzsche/pkg/file"
	"github.com/sebsvt/nietzsche/pkg/logging"
	"github.com/sebsvt/nietzsche/pkg/storage"
)

type Nietzsche interface {
	Start() (*StartResponse, error)
	Upload(filename string, content []byte) (*UploadResponse, error)
	Download(serverFileName string) ([]byte, error)
	Process(serverFileName string) (*int, error)
}

type nietzsche struct {
	storage storage.FileStorage
}

func NewNietzsche(storage storage.FileStorage) *nietzsche {
	return &nietzsche{
		storage: storage,
	}
}

func (n *nietzsche) Start() (*StartResponse, error) {
	res := &StartResponse{
		Server:           "https://nietzsche.sebsvt.com",
		Task:             "task_" + uuid.New().String(),
		RemainingCredits: rand.Intn(9999),
	}
	logging.Info(fmt.Sprintf("Start response task id: %+v", res.Task))
	return res, nil
}

// upload file via local file | upload file via url
func (n *nietzsche) Upload(filename string, content []byte) (*UploadResponse, error) {
	// extract extension from filename
	extension := filepath.Ext(filename)
	// validate extension
	if !file.IsValidExtension(filename, []string{".pdf", ".jpg", ".jpeg", ".png", ".gif"}) {
		return nil, errors.New("invalid file type")
	}

	// generate new filename for server filename
	serverFilename := uuid.New().String() + extension

	reader := bytes.NewReader(content)
	_, err := n.storage.Save(serverFilename, reader)
	if err != nil {
		return nil, err
	}
	logging.Info(fmt.Sprintf("Uploaded file to server: %+v", serverFilename))
	return &UploadResponse{
		ServerFileName: serverFilename,
	}, nil
}

func (n *nietzsche) Download(serverFileName string) ([]byte, error) {
	content, err := n.storage.Get(serverFileName)
	defer content.Close()
	if err != nil {
		return nil, err
	}
	contentBytes, err := io.ReadAll(content)
	if err != nil {
		return nil, err
	}
	return contentBytes, nil
}

func (n *nietzsche) Process(serverFileName string) (*int, error) {
	return nil, nil
}
