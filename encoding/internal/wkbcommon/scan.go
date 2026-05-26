package wkbcommon

import (
	"errors"

	"github.com/paulmach/orb"
)

var (
	// ErrUnsupportedDataType is returned by Scan methods when asked to scan
	// non []byte data from the database. This should never happen
	// if the driver is acting appropriately.
	ErrUnsupportedDataType = errors.New("wkbcommon: scan value must be []byte")

	// ErrNotWKB is returned when unmarshalling WKB and the data is not valid.
	ErrNotWKB = errors.New("wkbcommon: invalid data")

	// ErrNotWKBHeader is returned when unmarshalling first few bytes and there
	// is an issue.
	ErrNotWKBHeader = errors.New("wkbcommon: invalid header data")

	// ErrIncorrectGeometry is returned when unmarshalling WKB data into the wrong type.
	// For example, unmarshaling linestring data into a point.
	ErrIncorrectGeometry = errors.New("wkbcommon: incorrect geometry")

	// ErrUnsupportedGeometry is returned when geometry type is not supported by this lib.
	ErrUnsupportedGeometry = errors.New("wkbcommon: unsupported geometry")
)

// Scan will scan the input []byte data into a geometry.
// This could be into the orb geometry type pointer or, if nil,
// the scanner.Geometry attribute.
func Scan(g, d any) (orb.Geometry, int, bool, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), 0, false, nil
}

// go-pg will return ST_AsBinary(*) data as `\xhexencoded` which
// needs to be converted to true binary for further decoding.
// Code detects the \x prefix and then converts the rest from Hex to binary.

// also possible is just straight hex encoded.
// In this case the bo bit can be '0x00' or '0x01'

// ScanPoint takes binary wkb and decodes it into a point.
func ScanPoint(data []byte) (orb.Point, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0, nil
}

// ScanMultiPoint takes binary wkb and decodes it into a multi-point.
func ScanMultiPoint(data []byte) (orb.MultiPoint, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint), 0, nil
}

// ScanLineString takes binary wkb and decodes it into a line string.
func ScanLineString(data []byte) (orb.LineString, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), 0, nil
}

// ScanMultiLineString takes binary wkb and decodes it into a multi-line string.
func ScanMultiLineString(data []byte) (orb.MultiLineString, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString), 0, nil
}

// ScanPolygon takes binary wkb and decodes it into a polygon.
func ScanPolygon(data []byte) (orb.Polygon, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Polygon), 0, nil
}

// ScanMultiPolygon takes binary wkb and decodes it into a multi-polygon.
func ScanMultiPolygon(data []byte) (orb.MultiPolygon, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon), 0, nil
}

// ScanCollection takes binary wkb and decodes it into a collection.
func ScanCollection(data []byte) (orb.Collection, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Collection), 0, nil
}
