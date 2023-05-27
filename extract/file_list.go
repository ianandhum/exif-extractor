package extract

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ianandhum/exif-extractor/util"
)

type FileList struct {
	SourceDir          string
	IncludeHiddenFiles bool

	files []string
}

func (dirList *FileList) GetFiles() ([]string, error) {
	if dirList.files == nil {
		return nil, fmt.Errorf("files list is empty")
	}

	return dirList.files, nil
}

func (dirList *FileList) Populate() (err error) {
	return filepath.Walk(dirList.SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		_, fileName := filepath.Split(path)
		isHidden, err := util.IsHiddenFile(fileName)
		if err != nil {
			return nil
		}

		if !dirList.IncludeHiddenFiles && isHidden {
			return nil
		}

		dirList.files = append(dirList.files, path)
		return nil
	})
}
