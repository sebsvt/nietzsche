package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	CreateFile("test.txt")
	defer DeleteFile("test.txt")

	assert.True(t, FileExists("test.txt"))
}

func TestFolderExists(t *testing.T) {
	CreateDir("test")
	defer DeleteDir("test")

	assert.True(t, FolderExists("test"))
}

func TestCreateDir(t *testing.T) {
	CreateDir("test")
	defer DeleteDir("test")

	assert.NoError(t, CreateDir("test"))
}

func TestCreateFile(t *testing.T) {
	CreateFile("test.txt")
	defer DeleteFile("test.txt")

	assert.NoError(t, CreateFile("test.txt"))
}

func TestDeleteFile(t *testing.T) {
	CreateFile("test.txt")

	assert.NoError(t, DeleteFile("test.txt"))
}

func TestDeleteDir(t *testing.T) {
	CreateDir("test")

	assert.NoError(t, DeleteDir("test"))
}

func TestWriteFile(t *testing.T) {
	WriteFile("test.txt", []byte("Hello, World!"))
	defer DeleteFile("test.txt")

	assert.NoError(t, WriteFile("test.txt", []byte("Hello, World!")))
}
