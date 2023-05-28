package export

import (
	"fmt"
	"io"
)

type Exporter interface {
	ExportAsBytes(headers []string, data [][]string) ([]byte, error)
	ExportAsStream(headers []string, data [][]string, writer io.Writer) error
}

type ExporterType string

const (
	CSVExporterType ExporterType = "csv"
	HTMLExporerType ExporterType = "html"
)

func NewExporter(exporterType ExporterType) (Exporter, error) {
	switch exporterType {
	case CSVExporterType:
		return new(csvExporter), nil
	case HTMLExporerType:
		return new(htmlExporter), nil
	default:
		return nil, fmt.Errorf("unknown exporter type: %s", exporterType)
	}
}
