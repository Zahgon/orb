package wkt

import (
	"errors"
	"regexp"

	"github.com/paulmach/orb"
)

var (
	// ErrNotWKT is returned when unmarshalling WKT and the data is not valid.
	ErrNotWKT = errors.New("wkt: invalid data")

	// ErrIncorrectGeometry is returned when unmarshalling WKT data into the wrong type.
	// For example, unmarshaling linestring data into a point.
	ErrIncorrectGeometry = errors.New("wkt: incorrect geometry")

	// ErrUnsupportedGeometry is returned when geometry type is not supported by this lib.
	ErrUnsupportedGeometry = errors.New("wkt: unsupported geometry")

	doubleParen = regexp.MustCompile(`\)[\s|\t]*\)([\s|\t]*,[\s|\t]*)\([\s|\t]*\(`)
	singleParen = regexp.MustCompile(`\)([\s|\t]*,[\s|\t]*)\(`)
)

// UnmarshalPoint returns the point represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a point.
func UnmarshalPoint(s string) (orb.Point, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

func unmarshalPoint(s string) (orb.Point, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

// parsePoint parse point by (x y)
func parsePoint(s string) (p orb.Point, err error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

// UnmarshalMultiPoint returns the multi-point represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a multi-point.
func UnmarshalMultiPoint(s string) (orb.MultiPoint, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint), nil
}

func unmarshalMultiPoint(s string) (orb.MultiPoint, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint), nil
}

// UnmarshalLineString returns the linestring represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a linestring.
func UnmarshalLineString(s string) (orb.LineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

func unmarshalLineString(s string) (orb.LineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

// UnmarshalMultiLineString returns the multi-linestring represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a multi-linestring.
func UnmarshalMultiLineString(s string) (orb.MultiLineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString), nil
}

func unmarshalMultiLineString(s string) (orb.MultiLineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString), nil
}

// UnmarshalPolygon returns the polygon represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a polygon.
func UnmarshalPolygon(s string) (orb.Polygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.Polygon), nil
}

func unmarshalPolygon(s string) (orb.Polygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.Polygon), nil
}

// UnmarshalMultiPolygon returns the multi-polygon represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a multi-polygon.
func UnmarshalMultiPolygon(s string) (orb.MultiPolygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon), nil
}

func unmarshalMultiPolygon(s string) (orb.MultiPolygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon), nil
}

// UnmarshalCollection returns the geometry collection represented by the wkt string.
// Will return ErrIncorrectGeometry if the wkt is not a geometry collection.
func UnmarshalCollection(s string) (orb.Collection, error) {
	_ = "STUB: not implemented"
	return *new(orb.Collection), nil
}

func unmarshalCollection(s string) (orb.Collection, error) {
	_ = "STUB: not implemented"
	return *new(orb.Collection), nil
}

// just GEOMETRYCOLLECTION

// splitGeometryCollection split GEOMETRYCOLLECTION to more geometry
func splitGeometryCollection(s string) (r []string) { _ = "STUB: not implemented"; return nil }

// Unmarshal return a geometry by parsing the WKT string.
func Unmarshal(s string) (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

// splitByRegexpYield splits the input by the regexp. The first callback can
// be used to initialize an array with the size of the result, the second
// is the callback with the matches.
// We use a yield function because it was faster/used less memory than
// allocating an array of the results.
func splitByRegexpYield(s string, re *regexp.Regexp, set func(int), yield func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

// splitOnComma is optimized to split on the regex [\s|\t|\n]*,[\s|\t|\n]*
// i.e. comma with possible spaces on each side. e.g. '  ,  '
// We use a yield function because it was faster/used less memory than
// allocating an array of the results.
func splitOnComma(s string, yield func(s string) error) error {
	_ = "STUB: not implemented"
	// in WKT points are separated by commas, coordinates in points are separated by spaces
	// e.g. 1 2,3 4,5 6,7 81 2,5 4
	// we want to split this and find each point.
	return nil
}

// at is right after the previous space-comma-space match.
// once a space-comma-space match is found, we go from 'at' to the start
// of the match, that's the split that needs to be returned.

// the start of a space-comma-space section

// a space starts a section, we need to see a comma for it to be a valid section

// trimSpaceBrackets trim space and brackets
func trimSpaceBrackets(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func trimSpace(s string) string { _ = "STUB: not implemented"; return "" }

// gets the ToUpper case of the first 20 chars.
// This is to determine the type without doing a full strings.ToUpper
func upperPrefix(s string) []byte { _ = "STUB: not implemented"; return nil }

// copied here from strings.Cut so we don't require go1.18
func cut(s, sep string) (before, after string, found bool) {
	_ = "STUB: not implemented"
	return "", "", false
}
