// Package project defines projections to and from Mercator and WGS84
// along with helpers to apply them to orb geometry types.
package project

import "github.com/paulmach/orb"

// Geometry is a helper to project any geometry.
func Geometry(g orb.Geometry, proj orb.Projection) orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

// Point is a helper to project a point
func Point(p orb.Point, proj orb.Projection) orb.Point {
	_ = "STUB: not implemented"

	// MultiPoint is a helper to project an entire multi point.
	return *new(orb.Point)
}

func MultiPoint(mp orb.MultiPoint, proj orb.Projection) orb.MultiPoint {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint)
}

// LineString is a helper to project an entire line string.
func LineString(ls orb.LineString, proj orb.Projection) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

// MultiLineString is a helper to project an entire multi linestring.
func MultiLineString(mls orb.MultiLineString, proj orb.Projection) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// Ring is a helper to project an entire ring.
func Ring(r orb.Ring, proj orb.Projection) orb.Ring {
	_ = "STUB: not implemented"
	return *new(orb.Ring)
}

// Polygon is a helper to project an entire polygon.
func Polygon(p orb.Polygon, proj orb.Projection) orb.Polygon {
	_ = "STUB: not implemented"
	return *new(orb.Polygon)
}

// MultiPolygon is a helper to project an entire multi polygon.
func MultiPolygon(mp orb.MultiPolygon, proj orb.Projection) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// Collection is a helper to project a rectangle.
func Collection(c orb.Collection, proj orb.Projection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}

// Bound is a helper to project a rectangle.
func Bound(bound orb.Bound, proj orb.Projection) orb.Bound {
	_ = "STUB: not implemented"
	return *new(orb.Bound)
}
