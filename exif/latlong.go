package exif

// Reference: https://stackoverflow.com/a/1140335

func convertToDecimalDegrees(gpsDegrees GpsDegrees) float64 {
	decimalDegrees := gpsDegrees.Degrees + (gpsDegrees.Minutes / 60) + (gpsDegrees.Seconds / 3600)

	if gpsDegrees.Orientation == 'S' || gpsDegrees.Orientation == 'W' {
		decimalDegrees = -decimalDegrees
	}

	return decimalDegrees
}

func ConvertToLatLong(gpsInfo *GpsInfo) (float64, float64) {
	latitude := convertToDecimalDegrees(gpsInfo.Latitude)
	longitude := convertToDecimalDegrees(gpsInfo.Longitude)

	return latitude, longitude
}
