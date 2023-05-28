package exif

import (
	"fmt"
	"time"
)

// RawImageBytes represents entire image as []byte
type RawImageBytes = []byte

// RawExifBytes represents exif section of an image as []byte
type RawExifBytes = []byte

// Reader parses EXIF data from and image
type Reader interface {
	GetExifBlob(image RawImageBytes) (rawExif RawExifBytes, err error)
	GetGPSInfo(exifData RawExifBytes) (*GpsInfo, error)
}

// GpsDegrees Taken from: https://github.com/dsoprea/go-exif/blob/master/v3/gps.go#L21
type GpsDegrees struct {
	Orientation byte

	Degrees float64

	Minutes float64

	Seconds float64
}

// GpsInfo represents a GPS co-ordinate
type GpsInfo struct {
	Latitude, Longitude GpsDegrees
	Altitude            int
	Timestamp           time.Time
}

// ReaderType represents the exif reader library to be used
type ReaderType string

const (
	// GoExifLibrary parse EXIF info using go-exif go module
	GoExifLibrary ReaderType = "go-exif"
)

// NewExifReader creates a new instance of exif.Reader
func NewExifReader(parsingLib ReaderType) (Reader, error) {
	switch parsingLib {
	case GoExifLibrary:
		return new(goExifReader), nil
	default:
		return nil, fmt.Errorf("unknown exif reader type: %s", parsingLib)
	}
}
