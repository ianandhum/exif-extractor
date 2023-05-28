package exif_test

import (
	"testing"

	"github.com/ianandhum/exif-extractor/exif"
	"github.com/stretchr/testify/assert"
)

func TestConvertToLatLong(t *testing.T) {
	gpsInfo := &exif.GpsInfo{
		Latitude: exif.GpsDegrees{
			Degrees:     43,
			Minutes:     39,
			Seconds:     11.613,
			Orientation: 'N',
		},
		Longitude: exif.GpsDegrees{
			Degrees:     79,
			Minutes:     22,
			Seconds:     59.4618,
			Orientation: 'W',
		},
	}

	expectedLatitude := 43.65322583
	expectedLongitude := -79.38318383

	latitude, longitude := exif.ConvertToLatLong(gpsInfo)
	assert.InDelta(t, latitude, expectedLatitude, 0.00000001)
	assert.InDelta(t, longitude, expectedLongitude, 0.00000001)

}
