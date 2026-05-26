package geo

import (
	"github.com/paulmach/orb"
)

// NewBoundAroundPoint creates a new bound given a center point,
// and a distance from the center point in meters.
func NewBoundAroundPoint(center orb.Point, distance float64) orb.Bound {
	_ = "STUB: not implemented"
	return *new(orb.Bound)
}

// BoundPad expands the bound in all directions by the given amount of meters.
func BoundPad(b orb.Bound, meters float64) orb.Bound {
	_ = "STUB: not implemented"
	return *new(orb.Bound)
}

// BoundHeight returns the approximate height in meters.
func BoundHeight(b orb.Bound) float64 { _ = "STUB: not implemented"; return 0 }

// BoundWidth returns the approximate width in meters
// of the center of the bound.
func BoundWidth(b orb.Bound) float64 { _ = "STUB: not implemented"; return 0 }

// MinLatitude is the minimum possible latitude
var minLatitude = deg2rad(-90)

// MaxLatitude is the maximum possible latitude
var maxLatitude = deg2rad(90)

// MinLongitude is the minimum possible longitude
var minLongitude = deg2rad(-180)

// MaxLongitude is the maximum possible longitude
var maxLongitude = deg2rad(180)

func deg2rad(d float64) float64 { _ = "STUB: not implemented"; return 0 }

func rad2deg(r float64) float64 { _ = "STUB: not implemented"; return 0 }
