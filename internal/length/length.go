package length

import (
	"github.com/paulmach/orb"
)

// Length returns the length of the boundary of the geometry
// using 2d euclidean geometry.
func Length(g orb.Geometry, df orb.DistanceFunc) float64 { _ = "STUB: not implemented"; return 0 }

func lineStringLength(ls orb.LineString, df orb.DistanceFunc) float64 {
	_ = "STUB: not implemented"
	return 0
}

func polygonLength(p orb.Polygon, df orb.DistanceFunc) float64 { _ = "STUB: not implemented"; return 0 }
