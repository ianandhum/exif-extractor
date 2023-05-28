package export

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
)

type htmlExporter struct {
}

var htmlTemplate *template.Template

func getHTMLTemplate() *template.Template {
	if htmlTemplate == nil {
		parsedTpl, err := template.New("exif-export-template").Parse(htmlExportTemplate)
		if err != nil {
			panic(err)
		}

		htmlTemplate = parsedTpl
	}
	return htmlTemplate
}

func (exporter *htmlExporter) ExportAsBytes(headers []string, data [][]string) ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := exporter.ExportAsStream(headers, data, buffer); err != nil {
		return nil, fmt.Errorf("unable to export as html: %s", err)
	}

	return buffer.Bytes(), nil
}

func (exporter *htmlExporter) ExportAsStream(headers []string, data [][]string, writer io.Writer) error {

	htmlTpl := getHTMLTemplate()

	dataRows := templateInput{
		Rows:   data,
		Title:  "Html Report",
		Header: headers,
	}

	return htmlTpl.Execute(writer, dataRows)
}
