package export

import (
	"fmt"
	"io"
)

// Exporter write the given tabular data into a stream
type Exporter interface {
	ExportAsBytes(headers []string, data [][]string) ([]byte, error)
	ExportAsStream(headers []string, data [][]string, writer io.Writer) error
}

// ExporterType the impl type of the exporter
type ExporterType string

const (
	// CSVExporterType export in csv format
	CSVExporterType ExporterType = "csv"

	// HTMLExporerType export as html
	HTMLExporerType ExporterType = "html"
)

// NewExporter create a new exporter with the given export type
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
