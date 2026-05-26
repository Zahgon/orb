package clip

import (
	"github.com/paulmach/orb"
)

// Code based on https://github.com/mapbox/lineclip

// line will clip a line into a set of lines
// along the bounding box boundary.
func line(box orb.Bound, in orb.LineString, open bool) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// loops through all the intersection of the line and box.
// eg. across a corner could have two intersections.

// both points are in the box, accept

// segment went outside

// both on one side of the box.
// segment not part of the final result.

// A is outside, B is inside, clip edge

// B is outside, A is inside, clip edge

// new start is the old end

func push(out orb.MultiLineString, i int, p orb.Point) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// ring will clip the Ring into a smaller ring around the bounding box boundary.
func ring(box orb.Bound, in orb.Ring) orb.Ring { _ = "STUB: not implemented"; return *new(orb.Ring) }

// if we're not a nice closed ring, don't implicitly close it.

// if segment goes through the clip window, add an intersection

// swap back

// need to make sure our output is also closed.

// bitCode returns the point position relative to the bbox:
//
//	     left  mid  right
//	top  1001  1000  1010
//	mid  0001  0000  0010
//
// bottom  0101  0100  0110
func bitCode(b orb.Bound, p orb.Point) int { _ = "STUB: not implemented"; return 0 }

func bitCodeOpen(b orb.Bound, p orb.Point) int { _ = "STUB: not implemented"; return 0 }

// intersect a segment against one of the 4 lines that make up the bbox
func intersect(box orb.Bound, edge int, a, b orb.Point) orb.Point {
	_ = "STUB: not implemented"

	// top
	return *new(orb.Point)
}

// bottom

// right

// left
