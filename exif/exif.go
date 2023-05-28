package exif

import (
	"time"
)

type RawImageBytes = []byte

type RawExifBytes = []byte

type ExifReader interface {
	GetExifBlob(image *RawImageBytes) (rawExif *[]byte, err error)
	GetGPSInfo(exifData *RawExifBytes) (*GpsInfo, error)
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

const (
	GoExifLibrary = iota
)

func GetNewExifReader(parsingLib int) (ExifReader, error) {
	switch parsingLib {
	default:
		return new(GoExifReader), nil
	}
}
