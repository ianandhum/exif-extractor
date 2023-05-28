package export

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

type csvExporter struct {
}

func (exporter *csvExporter) ExportAsBytes(headers []string, data [][]string) ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := exporter.ExportAsStream(headers, data, buffer); err != nil {
		return nil, fmt.Errorf("unable to export as csv: %s", err)
	}

	return buffer.Bytes(), nil
}

func (exporter *csvExporter) ExportAsStream(headers []string, data [][]string, writer io.Writer) error {
	csvWriter := csv.NewWriter(writer)

	// header
	if len(headers) > 0 {
		if err := csvWriter.Write(headers); err != nil {
			return err
		}
	}

	// all rows
	if err := csvWriter.WriteAll(data); err != nil {
		return err
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
