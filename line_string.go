package orb

// LineString represents a set of points to be thought of as a polyline.
type LineString []Point

// GeoJSONType returns the GeoJSON type for the object.
func (ls LineString) GeoJSONType() string { _ = "STUB: not implemented"; return "" }

// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineString) Dimensions() int {
	_ = "STUB: not implemented"

	// Reverse will reverse the line string.
	// This is done inplace, ie. it modifies the original data.
	return 0
}

func (ls LineString) Reverse() { _ = "STUB: not implemented"; return }

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (ls LineString) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two line strings. Returns true if lengths are the same
// and all points are Equal.
func (ls LineString) Equal(lineString LineString) bool { _ = "STUB: not implemented"; return false }

// Clone returns a new copy of the line string.
func (ls LineString) Clone() LineString { _ = "STUB: not implemented"; return *new(LineString) }
