package extract

import (
	"fmt"
	"log"
	"os"

	"github.com/h2non/filetype"
	"github.com/ianandhum/exif-extractor/exif"
)

// GPSInfoFromDir return a map containing GPS EXIF information of all elibile image files in a directory
func GPSInfoFromDir(folderPath string, exifLib exif.ReaderType, includeHiddenFiles bool) (map[string]*exif.GpsInfo, error) {
	walkOptions := new(WalkOptions)
	walkOptions.SourceDir = folderPath
	walkOptions.IncludeHiddenFiles = includeHiddenFiles

	files, err := GetFilesInDir(walkOptions)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting file list: %s", err)
	}

	resultMap := map[string]*exif.GpsInfo{}

	for _, file := range files {
		gpsInfo, err := GPSInfoFromFile(file, exifLib)
		if err != nil {
			// TODO should this be fatal?
			log.Printf("WARN: unable to read gpsInfo: %s(path: %s)", err, file)
			continue
		}

		resultMap[file] = gpsInfo
	}

	return resultMap, nil
}

// GPSInfoFromFile return exif.GpsInfo from a given image file
func GPSInfoFromFile(filePath string, exifLib exif.ReaderType) (*exif.GpsInfo, error) {
	content, err := readImageFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error occured while extracting gps info: %s", err)
	}
	exifReader, err := exif.NewExifReader(exifLib)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting exif reader: %s", err)
	}

	return exifReader.GetGPSInfo(content)
}

func readImageFile(filePath string) (exif.RawImageBytes, error) {

	if !isImageFile(filePath) {
		return nil, fmt.Errorf("'%s' is not and image file", filePath)
	}

	imageBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error occured while reading file: %s", err)
	}

	return imageBytes, nil
}

func isImageFile(filePath string) bool {
	file, err := os.Open(filePath)

	if err != nil {
		return false
	}

	defer func() {
		file.Close()
	}()

	// only first 261 bytes are needed
	head := make([]byte, 261)
	file.Read(head)

	return filetype.IsImage(head)
}
