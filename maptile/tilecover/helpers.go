// Package tilecover computes the covering set of tiles for an orb.Geometry.
package tilecover

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/maptile"
)

// Geometry returns the covering set of tiles for the given geometry.
func Geometry(g orb.Geometry, z maptile.Zoom) (maptile.Set, error) {
	_ = "STUB: not implemented"
	return *new(maptile.Set), nil
}

// Point creates a tile cover for the point, i.e. just the tile
// containing the point.
func Point(ll orb.Point, z maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

// MultiPoint creates a tile cover for the set of points,
func MultiPoint(mp orb.MultiPoint, z maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

// Bound creates a tile cover for the bound. i.e. all the tiles
// that intersect the bound.
func Bound(b orb.Bound, z maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

// Collection returns the covering set of tiles for the
// geometry collection.
func Collection(c orb.Collection, z maptile.Zoom) (maptile.Set, error) {
	_ = "STUB: not implemented"
	return *new(maptile.Set), nil
}
