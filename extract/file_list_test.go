package extract_test

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/ianandhum/exif-extractor/extract"
	"github.com/stretchr/testify/assert"
)

const (
	imagesFolder      = "../testdata/images/"
	expectedFileCount = 7
)

func TestFileList(t *testing.T) {
	defer func() {
		if state := recover(); state != nil {
			err := state.(error)
			t.Fatalf("Test failure. %s", err)
		}
	}()

	cwd, err := os.Getwd()
	assert.NoError(t, err)

	t.Run("return correct number of files", func(t *testing.T) {

		path, err := filepath.Abs(path.Join(cwd, imagesFolder))
		assert.NoError(t, err)

		walkOptions := new(extract.WalkOptions)
		walkOptions.SourceDir = path
		walkOptions.IncludeHiddenFiles = false

		files, err := extract.GetFilesInDir(walkOptions)
		assert.NoError(t, err)
		assert.Equal(t, expectedFileCount, len(files))
	})
}
