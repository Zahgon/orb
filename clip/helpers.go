// Package clip is a library for clipping geometry to a bounding box.
package clip

import (
	"github.com/paulmach/orb"
)

// Geometry will clip the geometry to the bounding box using the
// correct functions for the type.
// This operation will modify the input of '1d or 2d geometry' by using as a
// scratch space so clone if necessary.
func Geometry(b orb.Bound, g orb.Geometry) orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

// Intersect check above

// MultiPoint returns a new set with the points outside the bound removed.
func MultiPoint(b orb.Bound, mp orb.MultiPoint) orb.MultiPoint {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint)
}

// LineString clips the linestring to the bounding box.
func LineString(b orb.Bound, ls orb.LineString, opts ...Option) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// MultiLineString clips the linestrings to the bounding box
// and returns a linestring union.
func MultiLineString(b orb.Bound, mls orb.MultiLineString, opts ...Option) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// Ring clips the ring to the bounding box and returns another ring.
// This operation will modify the input by using as a scratch space
// so clone if necessary.
func Ring(b orb.Bound, r orb.Ring) orb.Ring { _ = "STUB: not implemented"; return *new(orb.Ring) }

// Polygon clips the polygon to the bounding box excluding the inner rings
// if they do not intersect the bounding box.
// This operation will modify the input by using as a scratch space
// so clone if necessary.
func Polygon(b orb.Bound, p orb.Polygon) orb.Polygon {
	_ = "STUB: not implemented"
	return *new(orb.Polygon)
}

// MultiPolygon clips the multi polygon to the bounding box excluding
// any polygons if they don't intersect the bounding box.
// This operation will modify the input by using as a scratch space
// so clone if necessary.
func MultiPolygon(b orb.Bound, mp orb.MultiPolygon) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// Collection clips each element in the collection to the bounding box.
// It will exclude elements if they don't intersect the bounding box.
// This operation will modify the input of '2d geometry' by using as a
// scratch space so clone if necessary.
func Collection(b orb.Bound, c orb.Collection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}

// Bound intersects the two bounds. May result in an
// empty/degenerate bound.
func Bound(b, bound orb.Bound) orb.Bound { _ = "STUB: not implemented"; return *new(orb.Bound) }
