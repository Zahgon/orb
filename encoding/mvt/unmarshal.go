package mvt

import (
	"errors"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/mvt/vectortile"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/protoscan"
)

var ErrDataIsGZipped = errors.New("failed to unmarshal, data possibly gzipped")

// UnmarshalGzipped takes gzipped Mapbox Vector Tile (MVT) data and unzips it
// before decoding it into a set of layers, It does not project the coordinates.
func UnmarshalGzipped(data []byte) (Layers, error) {
	_ = "STUB: not implemented"
	return *new(Layers), nil
}

// Unmarshal takes Mapbox Vector Tile (MVT) data and converts into a
// set of layers, It does not project the coordinates.
func Unmarshal(data []byte) (Layers, error) { _ = "STUB: not implemented"; return *new(Layers), nil }

// decoder is here to reuse objects to save allocations/memory.
type decoder struct {
	keys     []string
	values   []any
	features [][]byte

	valMsg *protoscan.Message

	tags *protoscan.Iterator
	geom *protoscan.Iterator
}

func (d *decoder) Reset() { _ = "STUB: not implemented"; return }

func unmarshalTile(data []byte) (Layers, error) {
	_ = "STUB: not implemented"
	return *new(Layers), nil
}

func (d *decoder) Layer(msg *protoscan.Message) (*Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// version

// name

// feature

// keys

// values

// extent

func (d *decoder) Feature(msg *protoscan.Message) (*geojson.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// id

//tags, repeated packed

// geomtype

// geometry

func (d *decoder) Geometry(geomType vectortile.Tile_GeomType) (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

// A geomDecoder holds state for geometry decoding.
type geomDecoder struct {
	iter  *protoscan.Iterator
	count int
	used  int

	prev orb.Point
}

func (gd *geomDecoder) decodePoint() (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

func (gd *geomDecoder) decodeLine() (orb.LineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

func (gd *geomDecoder) decodeLineString() (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

func (gd *geomDecoder) decodePolygon() (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

// figure out if new polygon

func (gd *geomDecoder) cmdAndCount() (uint32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (gd *geomDecoder) NextPoint() (orb.Point, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

func (gd *geomDecoder) done() bool { _ = "STUB: not implemented"; return false }

func decodeValueMsg(msg *protoscan.Message) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Check if data is GZipped by reading the "magic bytes"
// Rarely this method can result in false positives
func dataIsGZipped(data []byte) bool { _ = "STUB: not implemented"; return false }
