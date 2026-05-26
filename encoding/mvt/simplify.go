package mvt

import (
	"github.com/paulmach/orb"
)

// Simplify will run all the geometry of all the layers through the provided simplifier.
func (ls Layers) Simplify(s orb.Simplifier) { _ = "STUB: not implemented"; return }

// Simplify will run the layer geometries through the simplifier.
func (l *Layer) Simplify(s orb.Simplifier) { _ = "STUB: not implemented"; return }

// RemoveEmpty will remove line strings shorter/smaller than the limits.
func (ls Layers) RemoveEmpty(lineLimit, areaLimit float64) { _ = "STUB: not implemented"; return }

// RemoveEmpty will remove line strings shorter/smaller than the limits.
func (l *Layer) RemoveEmpty(lineLimit, areaLimit float64) { _ = "STUB: not implemented"; return }

// point geometry

// line geometry
