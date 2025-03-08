package storage

import "io"

type FileStorage interface {
	Get(documentID string) (io.ReadCloser, error)
	Delete(documentID string) error
	URL(documentID string) (string, error)
	Save(documentID string, content io.Reader) (string, error)
	Close() error
}
