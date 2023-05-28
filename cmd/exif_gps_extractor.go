package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ianandhum/exif-extractor/exif"
	"github.com/ianandhum/exif-extractor/export"
	"github.com/ianandhum/exif-extractor/extract"
)

const exifReaderLib = exif.GoExifLibrary

type cmdOptions struct {
	sourceDir          string
	includeHiddenFiles bool

	outputFormat export.ExporterType
	outputPath   string
}

func exitWithUsage(message string) {
	fmt.Fprint(os.Stderr, message)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Usage: ")
	flag.PrintDefaults()
	os.Exit(1)
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func parseArgs() *cmdOptions {
	dir := flag.String("dir", "", "Path to image folder")
	includeHiddenFiles := flag.Bool("include-hidden-files", false, "Include hidden files in the directory")

	outputFormat := flag.String("format", "csv", "Output format, supported formats: csv,html")
	outputPath := flag.String("output", "-", "Destination path to store output (default write to standard output)")

	flag.Parse()

	sourceDir, err := filepath.Abs(filepath.Clean(*dir))
	if err != nil || !isDir(*dir) {
		exitWithUsage("Error: --dir argument should be a valid path")
	}

	if *outputPath != "-" && !isDir(filepath.Dir(*outputPath)) {
		exitWithUsage("Error, --output argument should have a valid dir path")
	}

	return &cmdOptions{
		sourceDir:          sourceDir,
		includeHiddenFiles: *includeHiddenFiles,

		outputFormat: export.ExporterType(*outputFormat),
		outputPath:   *outputPath,
	}
}

func flattenExifGpsInfo(exifInfo map[string]*exif.GpsInfo) [][]string {
	flatGpsInfo := [][]string{}
	for path, gpsInfo := range exifInfo {
		lat, lng := exif.ConvertToLatLong(gpsInfo)
		gpsArr := [3]string{path, strconv.FormatFloat(lat, 'f', 8, 64), strconv.FormatFloat(lng, 'f', 8, 64)}
		flatGpsInfo = append(flatGpsInfo, gpsArr[:])
	}

	return flatGpsInfo
}

func getOutputHeader() []string {
	return []string{"Image File Path", "Latitude", "Longitude"}
}

func main() {
	// Parse command-line arguments
	cmdOptions := parseArgs()

	exporter, err := export.NewExporter(cmdOptions.outputFormat)
	if err != nil {
		exitWithUsage(fmt.Sprintf("Invalid output format: %s", cmdOptions.outputFormat))
	}

	exifInfo, err := extract.GPSInfoFromDir(cmdOptions.sourceDir, exifReaderLib, cmdOptions.includeHiddenFiles)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured while reading exif info: %s", err)
		os.Exit(1)
	}

	var outputWriter io.Writer

	if cmdOptions.outputPath == "-" { // Write to standard output
		outputWriter = os.Stdout
	} else {
		outputFile, err := os.Create(cmdOptions.outputPath)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := outputFile.Close(); err != nil {
				panic(err)
			}
		}()

		outputWriter = outputFile
	}

	if err := exporter.ExportAsStream(getOutputHeader(), flattenExifGpsInfo(exifInfo), outputWriter); err != nil {
		fmt.Fprintf(os.Stderr, "error occured while writing result: %s", err)
	}

	if cmdOptions.outputPath != "-" {
		fmt.Fprintf(os.Stderr, "Extracted data successfully written to '%s'\n", cmdOptions.outputPath)
	}
}
