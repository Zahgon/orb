package wkt

import (
	"bytes"

	"github.com/paulmach/orb"
)

// Marshal returns a WKT representation of the geometry.
func Marshal(g orb.Geometry) []byte { _ = "STUB: not implemented"; return nil }

// MarshalString returns a WKT representation of the geometry as a string.
func MarshalString(g orb.Geometry) string { _ = "STUB: not implemented"; return "" }

func wkt(buf *bytes.Buffer, geom orb.Geometry) { _ = "STUB: not implemented"; return }

func writeLineString(buf *bytes.Buffer, ls orb.LineString) { _ = "STUB: not implemented"; return }
