//go:build !windows

package util

import "path/filepath"

func IsHiddenFile(path string) (bool, error) {
	_, fileName := filepath.Split(path)
	return fileName[0] == '.', nil
}
