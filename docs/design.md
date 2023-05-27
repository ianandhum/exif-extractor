# EXIF Extractor : Design document

This document encompasses the basic design for exif-extractor CLI.

**exif-extractor** scans image files and extracts exif data including location data to a csv file.
For parsing images and returning EXIF information stored in the image, go-exif library can be used.

The core system consists of the following parts:

1. Collecting a directory or file input to be parsed.
2. Reading Exif data from the image input.
3. Formatting the data, including geo-location data and file path.
4. Output processing: Converting the exif information about each file into a readable format as requested by the user.

Apart from this, the resulting binary should be a POSIX compliant executable as much as possible with POSIX cli arguments, stream inputs and output. Multi-platform support is very much encouraged.

