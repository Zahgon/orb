package planar

import (
	"github.com/paulmach/orb"
)

// RingContains returns true if the point is inside the ring.
// Points on the boundary are considered in.
func RingContains(r orb.Ring, point orb.Point) bool { _ = "STUB: not implemented"; return false }

// PolygonContains checks if the point is within the polygon.
// Points on the boundary are considered in.
func PolygonContains(p orb.Polygon, point orb.Point) bool { _ = "STUB: not implemented"; return false }

// MultiPolygonContains checks if the point is within the multi-polygon.
// Points on the boundary are considered in.
func MultiPolygonContains(mp orb.MultiPolygon, point orb.Point) bool {
	_ = "STUB: not implemented"
	return false
}

// Original implementation: http://rosettacode.org/wiki/Ray-casting_algorithm#Go
func rayIntersect(p, s, e orb.Point) (intersects, on bool) {
	_ = "STUB: not implemented"
	return false, false
}

// p == start

// vertical segment (s -> e)
// return true if within the line, check to see if start or end is greater.

// Move the y coordinate to deal with degenerate case

// matching the end point
