package exif_test

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ianandhum/exif-extractor/exif"
	"github.com/stretchr/testify/assert"
)

const (
	inValodEOFEXIFImageFilePath = "../testdata/images/exif-error.jpg"

	validXMPGPSImageFilePath  = "../testdata/images/more_images/dog.png"
	validEXIFGPSImageFilePath = "../testdata/images/more_images/wax-card.jpg"
	inValidGIFImageFilePath   = "../testdata/images/more_images/david.gif" // EXIF is not supported by GIF format
)

func TestExifReaderSelector(t *testing.T) {
	exifReader, err := exif.GetNewExifReader(exif.GoExifLibrary)
	assert.NoError(t, err)

	assert.True(t, reflect.TypeOf(exifReader).Name() != "GoExifReader")

}

func TestExifDataReader(t *testing.T) {
	defer func() {
		if state := recover(); state != nil {
			err := state.(error)
			t.Fatalf("Test failure. %s", err)
		}
	}()

	cwd, err := os.Getwd()
	assert.NoError(t, err)

	readImage := func(filePath string) *exif.RawImageBytes {
		imagePath, err := filepath.Abs(path.Join(cwd, filePath))
		assert.NoError(t, err)

		imageBytes, err := os.ReadFile(imagePath)
		assert.NoError(t, err, "Image read failed, fix the test case")

		return &imageBytes
	}

	exifReader := new(exif.GoExifReader)

	t.Run("fail on empty image input", func(t *testing.T) {
		emptyImage := exif.RawImageBytes{}
		_, err := exifReader.GetGPSInfo(&emptyImage)

		assert.Error(t, err)
	})

	t.Run("pass read image file corrupted exif data", func(t *testing.T) {

		imageBytes := readImage(inValodEOFEXIFImageFilePath)

		exifBytes, err := exifReader.GetExifBlob(imageBytes)
		assert.NoError(t, err)

		_, err = exifReader.GetIFDFromExifBytes(exifBytes)
		assert.Error(t, err)
	})

	t.Run("pass getting GPS info from image with XMP metadata", func(t *testing.T) {
		imageBytes := readImage(validXMPGPSImageFilePath)
		gpsInfo, err := exifReader.GetGPSInfo(imageBytes)

		assert.NoError(t, err)
		assert.NotEmpty(t, gpsInfo)
	})

	t.Run("pass getting GPS info from image with EXIF metadata", func(t *testing.T) {
		imageBytes := readImage(validEXIFGPSImageFilePath)
		gpsInfo, err := exifReader.GetGPSInfo(imageBytes)

		assert.NoError(t, err)
		assert.NotEmpty(t, gpsInfo)
	})

	t.Run("fail reading exif info from GIF file", func(t *testing.T) {

		imageBytes := readImage(inValidGIFImageFilePath)

		_, err := exifReader.GetExifBlob(imageBytes)
		assert.Error(t, err)
	})

	t.Run("fail reading GPS info from GIF file", func(t *testing.T) {

		imageBytes := readImage(inValidGIFImageFilePath)
		_, err = exifReader.GetGPSInfo(imageBytes)

		assert.Error(t, err)
	})
}
