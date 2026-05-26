// Package geo computes properties on geometries assuming they are lon/lat data.
package geo

import (
	"github.com/paulmach/orb"
)

// Area returns the area of the geometry on the earth.
func Area(g orb.Geometry) float64 { _ = "STUB: not implemented"; return 0 }

// SignedArea will return the signed area of the ring.
// Will return negative if the ring is in the clockwise direction.
// Will implicitly close the ring.
func SignedArea(r orb.Ring) float64 { _ = "STUB: not implemented"; return 0 }

func ringArea(r orb.Ring) float64 { _ = "STUB: not implemented"; return 0 }

// if not a closed ring, add an implicit calc for that last point.

// To support implicit closing of ring, replace references to
// the last point in r to the first 1.

// i = N-3

// i = N-2

// i = N-1

// i = 0 to N-3

func polygonArea(p orb.Polygon) float64 { _ = "STUB: not implemented"; return 0 }

func multiPolygonArea(mp orb.MultiPolygon) float64 { _ = "STUB: not implemented"; return 0 }

func collectionArea(c orb.Collection) float64 { _ = "STUB: not implemented"; return 0 }
