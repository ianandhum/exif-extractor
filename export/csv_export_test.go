package export_test

import (
	"bytes"
	"testing"

	"github.com/ianandhum/exif-extractor/export"
	"github.com/stretchr/testify/assert"
)

func TestCSVExport(t *testing.T) {

	exporter, err := export.NewExporter(export.CSVExporterType)
	assert.NoError(t, err)

	data := [][]string{
		{"/tmp/test.png", "77.12", "12.56"},
		{"/tmp/test2.png", "77.13", "12.52"},
		{"/tmp/test3.png", "77.10", "12.51"},
	}

	expected := "/tmp/test.png,77.12,12.56\n/tmp/test2.png,77.13,12.52\n/tmp/test3.png,77.10,12.51\n"
	expectedWithHeader := "File,Latitude,Longitude\n/tmp/test.png,77.12,12.56\n/tmp/test2.png,77.13,12.52\n/tmp/test3.png,77.10,12.51\n"

	t.Run("test csv export for simple document without header", func(t *testing.T) {

		csvData, err := exporter.ExportAsBytes([]string{}, data)

		assert.NoError(t, err)
		assert.EqualValues(t, expected, string(csvData))

	})

	t.Run("test csv export for simple document with header", func(t *testing.T) {

		header := []string{"File", "Latitude", "Longitude"}
		csvData, err := exporter.ExportAsBytes(header, data)

		assert.NoError(t, err)
		assert.EqualValues(t, expectedWithHeader, string(csvData))

	})

	t.Run("test csv export for simple document as stream", func(t *testing.T) {

		header := []string{"File", "Latitude", "Longitude"}
		buffer := new(bytes.Buffer)

		err := exporter.ExportAsStream(header, data, buffer)

		assert.NoError(t, err)
		assert.EqualValues(t, expectedWithHeader, buffer.String())

	})

	t.Run("test csv export for simple document as stream with partial calls", func(t *testing.T) {

		header := []string{"File", "Latitude", "Longitude"}
		buffer := new(bytes.Buffer)

		err := exporter.ExportAsStream(header, data[:2], buffer)
		assert.NoError(t, err)

		err = exporter.ExportAsStream([]string{}, data[2:], buffer)
		assert.NoError(t, err)

		assert.EqualValues(t, expectedWithHeader, buffer.String())

	})

}
