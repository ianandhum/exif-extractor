# GPS EXIF Data Extractor

**exif-extractor** scans image files and extracts exif data including location data to various file formats.

## Building

To build the binary locally, clone this repository and install go dependency modules 

* Go version 1.20 or greater. (not tested on other go versions)

Start by cloning the repository:

```
git clone https://github.com/ianandhum/exif-extractor.git
cd exif-extractor
```

Install dependencies by `go get` and build the binary

```
cd cmd
go get
go build exif_gps_extractor.go
```

## Usage: 

To run GPS EXIF data extractor, execute `exif_gps_extractor` binary with the necessary arguments
<br/>

```
  -dir string (required)
        Path to image folder
  -format string
        Output format, supported formats: csv,html (default "csv")
  -include-hidden-files
        Include hidden files in the directory
  -output string
        Destination path to store output (default write to standard output) (default "-")
```