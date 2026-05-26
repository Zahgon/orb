// Package maptile defines a Tile type and methods to work with
// web map projected tile data.
package maptile

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// Tiles is a set of tiles, later we can add methods to this.
type Tiles []Tile

// ToFeatureCollection converts the tiles into a feature collection.
// This method is mostly useful for debugging output.
func (ts Tiles) ToFeatureCollection() *geojson.FeatureCollection {
	_ = "STUB: not implemented"
	return nil
}

// Tile is an x, y, z web mercator tile.
type Tile struct {
	X, Y uint32
	Z    Zoom
}

// A Zoom is a strict type for a tile zoom level.
type Zoom uint32

// New creates a new tile with the given coordinates.
func New(x, y uint32, z Zoom) Tile {
	_ = "STUB: not implemented"
	return *

	// At creates a tile for the point at the given zoom.
	// Will create a valid tile for the zoom. Points outside
	// the range lat [-85.0511, 85.0511] will be snapped to the
	// max or min tile as appropriate.
	new(Tile)
}

func At(ll orb.Point, z Zoom) Tile { _ = "STUB: not implemented"; return *new(Tile) }

// FromQuadkey creates the tile from the quadkey.
func FromQuadkey(k uint64, z Zoom) Tile { _ = "STUB: not implemented"; return *new(Tile) }

// Valid returns if the tile's x/y are within the range for the tile's zoom.
func (t Tile) Valid() bool { _ = "STUB: not implemented"; return false }

// Bound returns the geo bound for the tile.
// An optional tileBuffer parameter can be passes to create a buffer
// around the bound in tile dimension. e.g. a tileBuffer of 1 would create
// a bound 9x the size of the tile, centered around the provided tile.
func (t Tile) Bound(tileBuffer ...float64) orb.Bound {
	_ = "STUB: not implemented"
	return *new(orb.Bound)
}

// Center returns the center of the tile.
func (t Tile) Center() orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

// Contains returns if the given tile is fully contained (or equal to) the give tile.
func (t Tile) Contains(tile Tile) bool { _ = "STUB: not implemented"; return false }

// Parent returns the parent of the tile.
func (t Tile) Parent() Tile { _ = "STUB: not implemented"; return *new(Tile) }

// Fraction returns the precise tile fraction at the given zoom.
// Will return 2^zoom-1 if the point is below 85.0511 S.
func Fraction(ll orb.Point, z Zoom) orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

// bound it because we have a top of the world problem

// SharedParent returns the tile that contains both the tiles.
func (t Tile) SharedParent(tile Tile) Tile {
	_ = "STUB: not implemented"
	// bring both tiles to the lowest zoom.
	return *new(Tile)
}

// go version < 1.9
// bit package usage was about 10% faster
//
// TODO: use build flags to support older versions of go.
//
// move from most significant to least until there isn't a match.
// for i := t.Z - 1; i >= 0; i-- {
// 	if t.X&(1<<i) != tile.X&(1<<i) ||
// 		t.Y&(1<<i) != tile.Y&(1<<i) {
// 		return Tile{
// 			t.X >> (i + 1),
// 			t.Y >> (i + 1),
// 			t.Z - (i + 1),
// 		}
// 	}
// }
//
// if we reach here the tiles are the same, which was checked above.
// panic("unreachable")

// bits different for x and y

// max of xc, yc

// Children returns the 4 children of the tile.
func (t Tile) Children() Tiles { _ = "STUB: not implemented"; return *new(Tiles) }

// ChildrenInZoomRange returns all the children tiles of tile from ranges [zoomStart, zoomEnd], both ends inclusive.
func ChildrenInZoomRange(tile Tile, zoomStart, zoomEnd Zoom) Tiles {
	_ = "STUB: not implemented"
	//nolint:staticcheck // clearer this way
	return *new(Tiles)
}

//nolint:staticcheck // clearer this way

// Siblings returns the 4 tiles that share this tile's parent.
func (t Tile) Siblings() Tiles { _ = "STUB: not implemented"; return *new(Tiles) }

// Quadkey returns the quad key for the tile.
func (t Tile) Quadkey() uint64 { _ = "STUB: not implemented"; return 0 }

// Range returns the min and max tile "range" to cover the tile
// at the given zoom.
func (t Tile) Range(z Zoom) (min, max Tile) {
	_ = "STUB: not implemented"
	return *new(Tile), *new(Tile)
}

func (t Tile) toZoom(z Zoom) Tile { _ = "STUB: not implemented"; return *new(Tile) }
