package orb

var emptyBound = Bound{Min: Point{1, 1}, Max: Point{-1, -1}}

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
//
//	orb.MultiPoint{p1, p2}.Bound()
type Bound struct {
	Min, Max Point
}

// GeoJSONType returns the GeoJSON type for the object.
func (b Bound) GeoJSONType() string {
	_ = "STUB: not implemented"

	// Dimensions returns 2 because a Bound is a 2d object.
	return ""
}

func (b Bound) Dimensions() int {
	_ = "STUB: not implemented"

	// ToPolygon converts the bound into a Polygon object.
	return 0
}

func (b Bound) ToPolygon() Polygon { _ = "STUB: not implemented"; return *new(Polygon) }

// ToRing converts the bound into a loop defined
// by the boundary of the box.
func (b Bound) ToRing() Ring { _ = "STUB: not implemented"; return *new(Ring) }

// Extend grows the bound to include the new point.
func (b Bound) Extend(point Point) Bound {
	_ = "STUB: not implemented"
	// already included, no big deal
	return *new(Bound)
}

// Union extends this bound to contain the union of this and the given bound.
func (b Bound) Union(other Bound) Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(point Point) bool { _ = "STUB: not implemented"; return false }

// Intersects determines if two bounds intersect.
// Returns true if they are touching.
func (b Bound) Intersects(bound Bound) bool { _ = "STUB: not implemented"; return false }

// Pad extends the bound in all directions by the given value.
func (b Bound) Pad(d float64) Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Center returns the center of the bounds by "averaging" the x and y coords.
func (b Bound) Center() Point { _ = "STUB: not implemented"; return *new(Point) }

// Top returns the top of the bound.
func (b Bound) Top() float64 {
	_ = "STUB: not implemented"

	// Bottom returns the bottom of the bound.
	return 0
}

func (b Bound) Bottom() float64 {
	_ = "STUB: not implemented"

	// Right returns the right of the bound.
	return 0
}

func (b Bound) Right() float64 {
	_ = "STUB: not implemented"

	// Left returns the left of the bound.
	return 0
}

func (b Bound) Left() float64 {
	_ = "STUB: not implemented"

	// LeftTop returns the upper left point of the bound.
	return 0
}

func (b Bound) LeftTop() Point { _ = "STUB: not implemented"; return *new(Point) }

// RightBottom return the lower right point of the bound.
func (b Bound) RightBottom() Point { _ = "STUB: not implemented"; return *new(Point) }

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsZero return true if the bound includes just null island.
func (b Bound) IsZero() bool { _ = "STUB: not implemented"; return false }

// Bound returns the same bound.
func (b Bound) Bound() Bound {
	_ = "STUB: not implemented"

	// Equal returns if two bounds are equal.
	return *new(Bound)
}

func (b Bound) Equal(c Bound) bool { _ = "STUB: not implemented"; return false }
