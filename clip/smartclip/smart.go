// Package smartclip performs a more advanced clipping algorithm so
// it can deal with correctly oriented open rings and polygon.
package smartclip

import (
	"github.com/paulmach/orb"
)

// Geometry will do a smart more involved clipping and wrapping of the geometry.
// It will return simple OGC geometries. Rings that are NOT closed AND have an
// endpoint in the bound will be implicitly closed.
func Geometry(box orb.Bound, g orb.Geometry, o orb.Orientation) orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

// Ring will smart clip a ring to the boundary. This may result multiple rings so
// a multipolygon is possible. Rings that are NOT closed AND have an endpoint in
// the bound will be implicitly closed.
func Ring(box orb.Bound, r orb.Ring, o orb.Orientation) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// nothing was clipped

// everything outside bound

// everything inside bound

// in a well defined ring there will be no closed sections

// Polygon will smart clip a polygon to the bound.
// Rings that are NOT closed AND have an endpoint in the bound will be
// implicitly closed.
func Polygon(box orb.Bound, p orb.Polygon, o orb.Orientation) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// nothing was clipped

// everything outside bound

// everything inside bound

// MultiPolygon will smart clip a multipolygon to the bound.
// Rings that are NOT closed AND have an endpoint in the bound will be
// implicitly closed.
func MultiPolygon(box orb.Bound, mp orb.MultiPolygon, o orb.Orientation) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// outer rings

// nothing was clipped

// everything outside bound

// everything inside bound

// inner rings

// smart wrap everything that touches the edges

// clipRings will take a set of rings and clip them to the boundary.
// It returns the open lineStrings with endpoints on the boundary and
// the closed interior rings.
func clipRings(box orb.Bound, rings []orb.Ring) (open []orb.LineString, closed []orb.Ring) {
	_ = "STUB: not implemented"
	return nil, nil
}

// outside of bound

// if the input was a closed ring where the endpoints were within the bound,
// then join the sections.

// This operation is O(n^2), however, n is the number of segments, not edges
// so I think it's manageable.

// endpoint must be within the bound to try join

// closed ring, so completely inside bound
// unless it touches a boundary

type endpoint struct {
	Point    orb.Point
	Start    bool
	Used     bool
	Side     uint8
	Index    int
	OtherEnd int
}

func (e *endpoint) Before(mls []orb.LineString) orb.Point {
	_ = "STUB: not implemented"
	return *new(orb.Point)
}

var emptyTwoRing = orb.Ring{{}, {}}

// smartWrap takes the open lineStrings with endpoints on the boundary and
// connects them correctly.
func smartWrap(box orb.Bound, input []orb.LineString, o orb.Orientation) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// start

// end

// this operation is O(n^2). Technically we could use a linked list
// and remove points instead of marking them as "used".
// However since n is 2x the number of segments I think we're okay.

// previous was end, connect to this start

// loop complete!!

// start over looking for unused endpoints

const notOnSide = 0xFF

// :    4
// :   +-+
// : 1 | | 3
// :   +-+
// :    2
func pointSide(b orb.Bound, p orb.Point) uint8 { _ = "STUB: not implemented"; return 0 }

type sortableEndpoints struct {
	mls []orb.LineString
	eps []*endpoint
}

func (e *sortableEndpoints) Len() int {
	_ = "STUB: not implemented"

	// Less sorts the points around the bound.
	// First comparing what side it's on and then the actual point to determine the order.
	// If two points are the same, we sort by the edge attached to the point so lines that are
	// "above" are shorted first.
	return 0
}

func (e *sortableEndpoints) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (e *sortableEndpoints) Swap(i, j int) { _ = "STUB: not implemented"; return }

// addToMultiPolygon does a lookup to see which polygon the ring intersects.
// This should work fine if the input is well formed.
func addToMultiPolygon(mp orb.MultiPolygon, ring orb.Ring) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// ring is not in any polygons?
// skip it, TODO: is this correct?
// If input is well formed, I think it is. If it isn't, ¯\_(ツ)_/¯

func polygonContains(outer orb.Ring, r orb.Ring) bool { _ = "STUB: not implemented"; return false }
