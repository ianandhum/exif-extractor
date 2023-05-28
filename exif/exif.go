package exif

import (
	"fmt"
	"time"
)

type RawImageBytes = []byte

type RawExifBytes = []byte

type ExifReader interface {
	GetExifBlob(image RawImageBytes) (rawExif RawExifBytes, err error)
	GetGPSInfo(exifData RawExifBytes) (*GpsInfo, error)
}

// Taken from: https://github.com/dsoprea/go-exif/blob/master/v3/gps.go#L21
type GpsDegrees struct {
	Orientation byte

	Degrees float64

	Minutes float64

	Seconds float64
}

type GpsInfo struct {
	Latitude, Longitude GpsDegrees
	Altitude            int
	Timestamp           time.Time
}

type ExifReaderType string

const (
	GoExifLibrary ExifReaderType = "go-exif"
)

func NewExifReader(parsingLib ExifReaderType) (ExifReader, error) {
	switch parsingLib {
	case GoExifLibrary:
		return new(goExifReader), nil
	default:
		return nil, fmt.Errorf("unknown exif reader type: %s", parsingLib)
	}
}
