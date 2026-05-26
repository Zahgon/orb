package tilecover

import (
	"errors"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/maptile"
)

// ErrUnevenIntersections can be returned when clipping polygons
// and there are issues with the geometries, like the rings are not closed.
var ErrUnevenIntersections = errors.New("tilecover: uneven intersections, ring not closed?")

// Ring creates a tile cover for the ring.
func Ring(r orb.Ring, z maptile.Zoom) (maptile.Set, error) {
	_ = "STUB: not implemented"
	return *new(maptile.Set), nil
}

// Polygon creates a tile cover for the polygon.
func Polygon(p orb.Polygon, z maptile.Zoom) (maptile.Set, error) {
	_ = "STUB: not implemented"
	return *new(maptile.Set), nil
}

// MultiPolygon creates a tile cover for the multi-polygon.
func MultiPolygon(mp orb.MultiPolygon, z maptile.Zoom) (maptile.Set, error) {
	_ = "STUB: not implemented"
	return *new(maptile.Set), nil
}

func polygon(set maptile.Set, p orb.Polygon, zoom maptile.Zoom) error {
	_ = "STUB: not implemented"
	return nil
}

// add intersection if it's not local extremum or duplicate
// not local minimum
// not local maximum

// sort by y, then x

// fill tiles between pairs of intersections
