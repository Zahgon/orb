package mvt

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/mvt/vectortile"
)

const (
	moveTo    = 1
	lineTo    = 2
	closePath = 7
)

func encodeGeometry(g orb.Geometry) (vectortile.Tile_GeomType, []uint32, error) {
	_ = "STUB: not implemented"
	return *new(vectortile.Tile_GeomType), nil, nil
}

type geomEncoder struct {
	prevX, prevY int32
	Data         []uint32
}

func newGeomEncoder(l int) *geomEncoder { _ = "STUB: not implemented"; return nil }

func (ge *geomEncoder) MoveTo(points []orb.Point) { _ = "STUB: not implemented"; return }

func (ge *geomEncoder) LineTo(points []orb.Point) { _ = "STUB: not implemented"; return }

func (ge *geomEncoder) addPoints(points []orb.Point) { _ = "STUB: not implemented"; return }

func (ge *geomEncoder) ClosePath() { _ = "STUB: not implemented"; return }

type keyValueEncoder struct {
	Keys   []string
	keyMap map[string]uint32

	Values   []*vectortile.Tile_Value
	valueMap map[any]uint32

	keySortBuffer []string
}

func newKeyValueEncoder() *keyValueEncoder { _ = "STUB: not implemented"; return nil }

func (kve *keyValueEncoder) Key(s string) uint32 { _ = "STUB: not implemented"; return 0 }

func (kve *keyValueEncoder) Value(v any) (uint32, error) {
	_ = "STUB: not implemented"
	// If a type is not comparable we can't figure out uniqueness in the hash,
	// we also can't encode it into a vectortile.Tile_Value.
	// So we encoded it as a json string, which is what other encoders
	// also do.
	return 0, nil
}

func encodeValue(v any) (*vectortile.Tile_Value, error) { _ = "STUB: not implemented"; return nil, nil }

func decodeValue(v *vectortile.Tile_Value) any { _ = "STUB: not implemented"; return *new(any) }

// functions to estimate encoded length

func elMLS(mls orb.MultiLineString) int { _ = "STUB: not implemented"; return 0 }

func elP(p orb.Polygon) int { _ = "STUB: not implemented"; return 0 }

func elMP(mp orb.MultiPolygon) int { _ = "STUB: not implemented"; return 0 }

func unzigzag(v uint32) float64 { _ = "STUB: not implemented"; return 0 }
