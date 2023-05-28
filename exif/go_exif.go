package exif

import (
	"fmt"

	goExif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type goExifReader struct {
}

func (reader *goExifReader) GetExifBlob(image RawImageBytes) (rawExif RawExifBytes, err error) {
	exifData, err := goExif.SearchAndExtractExif(image)
	if err != nil {
		if err == goExif.ErrNoExif {
			return nil, fmt.Errorf("no EXIF data found: %s", err)
		}
		return nil, fmt.Errorf("unable to read exif data from the image: %s", err)
	}
	return exifData, nil
}

func (reader *goExifReader) GetGPSInfo(imageBytes RawImageBytes) (*GpsInfo, error) {
	exifData, err := reader.GetExifBlob(imageBytes)
	if err != nil {
		return nil, fmt.Errorf("error while reading exif data: %s", err)
	}

	ifdIndex, err := reader.GetIFDFromExifBytes(exifData)
	if err != nil {
		return nil, fmt.Errorf("error while parsing exif data: %s", err)
	}

	gpsInfo, err := reader.getGPSInfoFromIFD(ifdIndex)
	return gpsInfo, err
}

func (reader *goExifReader) GetIFDFromExifBytes(exifBytes RawExifBytes) (*goExif.IfdIndex, error) {

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, fmt.Errorf("unable to create new IFD mapping: %s", err)
	}

	ti := goExif.NewTagIndex()
	_, index, err := goExif.Collect(im, ti, exifBytes)

	if err != nil {
		return nil, fmt.Errorf("unable to collect IFD tags : %s", err)
	}
	return &index, nil
}

func (reader *goExifReader) getGPSInfoFromIFD(ifdIndex *goExif.IfdIndex) (*GpsInfo, error) {

	// Get the GPS tag from exit data
	ifd, err := ifdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return nil, fmt.Errorf("unable to collect GPS IFD tag : %s", err)
	}
	gi, err := ifd.GpsInfo()
	if err != nil {
		return nil, fmt.Errorf("unable to get GPSInfo from IFD : %s", err)
	}

	var giResult *GpsInfo = &GpsInfo{
		Latitude:  GpsDegrees(gi.Latitude),
		Longitude: GpsDegrees(gi.Longitude),
		Altitude:  gi.Altitude,
		Timestamp: gi.Timestamp,
	}

	return giResult, nil
}
