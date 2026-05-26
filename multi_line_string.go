package orb

// MultiLineString is a set of polylines.
type MultiLineString []LineString

// GeoJSONType returns the GeoJSON type for the object.
func (mls MultiLineString) GeoJSONType() string { _ = "STUB: not implemented"; return "" }

// Dimensions returns 1 because a MultiLineString is a 2d object.
func (mls MultiLineString) Dimensions() int {
	_ = "STUB: not implemented"

	// Bound returns a bound around all the line strings.
	return 0
}

func (mls MultiLineString) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two multi line strings. Returns true if lengths are the same
// and all points are Equal.
func (mls MultiLineString) Equal(multiLineString MultiLineString) bool {
	_ = "STUB: not implemented"
	return false
}

// Clone returns a new deep copy of the multi line string.
func (mls MultiLineString) Clone() MultiLineString {
	_ = "STUB: not implemented"
	return *new(MultiLineString)
}
