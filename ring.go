package orb

// Ring represents a set of ring on the earth.
type Ring LineString

// GeoJSONType returns the GeoJSON type for the object.
func (r Ring) GeoJSONType() string {
	_ = "STUB: not implemented"

	// Dimensions returns 2 because a Ring is a 2d object.
	return ""
}

func (r Ring) Dimensions() int {
	_ = "STUB: not implemented"

	// Closed will return true if the ring is a real ring.
	// ie. 4+ points and the first and last points match.
	// NOTE: this will not check for self-intersection.
	return 0
}

func (r Ring) Closed() bool { _ = "STUB: not implemented"; return false }

// Reverse changes the direction of the ring.
// This is done inplace, ie. it modifies the original data.
func (r Ring) Reverse() { _ = "STUB: not implemented"; return }

// Bound returns a rect around the ring. Uses rectangular coordinates.
func (r Ring) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Orientation returns 1 if the ring is in counter-clockwise order,
// return -1 if the ring is the clockwise order and 0 if the ring is
// degenerate and had no area.
func (r Ring) Orientation() Orientation {
	_ = "STUB: not implemented"

	// This is a fast planar area computation, which is okay for this use.
	// implicitly move everything to near the origin to help with roundoff
	return *new(Orientation)
}

// degenerate case, no area

// Equal compares two rings. Returns true if lengths are the same
// and all points are Equal.
func (r Ring) Equal(ring Ring) bool { _ = "STUB: not implemented"; return false }

// Clone returns a new copy of the ring.
func (r Ring) Clone() Ring { _ = "STUB: not implemented"; return *new(Ring) }
