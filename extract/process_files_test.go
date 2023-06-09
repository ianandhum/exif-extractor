package extract_test

import (
	"testing"

	"github.com/ianandhum/exif-extractor/exif"
	"github.com/ianandhum/exif-extractor/extract"
	"github.com/stretchr/testify/assert"
)

const (
	imagesDir = "../testdata/images/"

	validEXIFGPSImageFilePath = "../testdata/images/more_images/wax-card.jpg"
	inValidGIFImageFilePath   = "../testdata/images/more_images/david.gif" // EXIF is not supported by GIF format
)

const (
	expectedGpsDirMapLength = 4
)

func TestFileProcessing(t *testing.T) {
	defer func() {
		if state := recover(); state != nil {
			err := state.(error)
			t.Fatalf("Test failure. %s", err)
		}
	}()

	t.Run("extract GPS info from a valid file", func(t *testing.T) {
		gpsInfo, err := extract.GPSInfoFromFile(validEXIFGPSImageFilePath, exif.GoExifLibrary)

		assert.NoError(t, err)
		assert.NotEmpty(t, gpsInfo)
	})

	t.Run("fail GPS info extraction on gif file", func(t *testing.T) {
		_, err := extract.GPSInfoFromFile(inValidGIFImageFilePath, exif.GoExifLibrary)

		assert.Error(t, err)
	})

	t.Run("extract gps info from image files in a directory", func(t *testing.T) {
		gpsInfoMap, err := extract.GPSInfoFromDir(imagesDir, exif.GoExifLibrary, false)

		assert.NoError(t, err)
		assert.Equal(t, expectedGpsDirMapLength, len(gpsInfoMap))
	})

}
