package tilecover

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/maptile"
)

// LineString creates a tile cover for the line string.
func LineString(ls orb.LineString, z maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

// MultiLineString creates a tile cover for the line strings.
func MultiLineString(mls orb.MultiLineString, z maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

func line(
	set maptile.Set,
	line orb.LineString,
	zoom maptile.Zoom,
	ring [][2]uint32,
) [][2]uint32 {
	_ = "STUB: not implemented"
	return nil
}
