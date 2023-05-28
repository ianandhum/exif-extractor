package extract

import (
	"fmt"
	"log"
	"os"

	"github.com/h2non/filetype"
	"github.com/ianandhum/exif-extractor/exif"
)

func ExtractGPSInfoFromDir(folderPath string, exifLib exif.ExifReaderType, includeHiddenFiles bool) (map[string]*exif.GpsInfo, error) {
	walkOptions := new(WalkOptions)
	walkOptions.SourceDir = folderPath
	walkOptions.IncludeHiddenFiles = includeHiddenFiles

	files, err := GetFilesInDir(walkOptions)
	if err != nil {
		return nil, fmt.Errorf("error occured while getting file list: %s", err)
	}

	resultMap := map[string]*exif.GpsInfo{}

	for _, file := range files {
		gpsInfo, err := ExtractGPSInfoFromFile(file, exifLib)
		if err != nil {
			// TODO should this be fatal?
			log.Printf("WARN: unable to read gpsInfo: %s(path: %s)", err, file)
			continue
		}

		resultMap[file] = gpsInfo
	}

	return resultMap, nil
}

func ExtractGPSInfoFromFile(filePath string, exifLib exif.ExifReaderType) (*exif.GpsInfo, error) {
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

	// only first 261 bytes are needed
	head := make([]byte, 261)
	file.Read(head)

	return filetype.IsImage(head)
}
