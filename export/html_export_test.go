package export_test

import (
	"bytes"
	"testing"

	"github.com/ianandhum/exif-extractor/export"
	"github.com/stretchr/testify/assert"
)

func TestHTMLExporter(t *testing.T) {

	exporter, err := export.NewExporter(export.HTMLExporerType)
	assert.NoError(t, err)

	data := [][]string{
		{"/tmp/test.png", "77.12", "12.56"},
		{"/tmp/test2.png", "77.13", "12.52"},
		{"/tmp/test3.png", "77.10", "12.51"},
	}

	header := []string{"File", "Latitude", "Longitude"}

	expectedWithoutHeader := "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset='utf-8'>\n    <meta http-equiv='X-UA-Compatible' content='IE=edge'>\n    <title>Html Report</title>\n    <meta name='viewport' content='width=device-width, initial-scale=1'>\n</head>\n<body>\n\t<table>\n\t<tr>\n\t\t\n\t</tr>\n\t<tbody>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.12</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.56</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test2.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.13</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.52</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test3.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.10</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.51</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t</tbody>\n\t</table>\n</body>\n</html>\n"
	expectedWithHeader := "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset='utf-8'>\n    <meta http-equiv='X-UA-Compatible' content='IE=edge'>\n    <title>Html Report</title>\n    <meta name='viewport' content='width=device-width, initial-scale=1'>\n</head>\n<body>\n\t<table>\n\t<tr>\n\t\t\n\t\t\t<th>File</th>\n\t\t\n\t\t\t<th>Latitude</th>\n\t\t\n\t\t\t<th>Longitude</th>\n\t\t\n\t</tr>\n\t<tbody>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.12</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.56</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test2.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.13</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.52</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t\t\t<tr>\n\t\t\t\t\n\t\t\t\t\t<td>/tmp/test3.png</td>\n\t\t\t\t\n\t\t\t\t\t<td>77.10</td>\n\t\t\t\t\n\t\t\t\t\t<td>12.51</td>\n\t\t\t\t\n\t\t\t</tr>\n\t\t\n\t</tbody>\n\t</table>\n</body>\n</html>\n"

	t.Run("test html export for simple document without header", func(t *testing.T) {

		csvData, err := exporter.ExportAsBytes([]string{}, data)

		assert.NoError(t, err)
		assert.EqualValues(t, expectedWithoutHeader, string(csvData))

	})

	t.Run("test html export for simple document with header", func(t *testing.T) {

		csvData, err := exporter.ExportAsBytes(header, data)

		assert.NoError(t, err)
		assert.EqualValues(t, expectedWithHeader, string(csvData))

	})

	t.Run("test html export for simple document as stream", func(t *testing.T) {

		buffer := new(bytes.Buffer)

		err := exporter.ExportAsStream(header, data, buffer)

		assert.NoError(t, err)
		assert.EqualValues(t, expectedWithHeader, buffer.String())

	})
}
