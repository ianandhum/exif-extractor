package extract

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ianandhum/exif-extractor/util"
)

type WalkOptions struct {
	SourceDir          string
	IncludeHiddenFiles bool
}

func GetFilesInDir(walkOptions *WalkOptions) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(walkOptions.SourceDir, func(path string, info os.FileInfo, err error) error {
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

		if !walkOptions.IncludeHiddenFiles && isHidden {
			return nil
		}

		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to get file list: %s", err)
	}

	return fileList, nil

}
