package mvt

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/mvt/vectortile"
	"github.com/paulmach/orb/geojson"
)

// MarshalGzipped will marshal the layers into Mapbox Vector Tile format
// and gzip the result. A lot of times MVT data is gzipped at rest,
// e.g. in a mbtiles file.
func MarshalGzipped(layers Layers) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalToVectorTile will take a set of layers and encode them into a Mapbox Vector Tile proto structure.
// Features that have a nil geometry, for some reason, will be skipped and not included.
func MarshalToVectorTile(layers Layers) (*vectortile.Tile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marshal will take a set of layers and encode them into a Mapbox Vector Tile format.
// Features that have a nil geometry, for some reason, will be skipped and not included.
func Marshal(layers Layers) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func addFeature(layer *vectortile.Tile_Layer, kve *keyValueEncoder, f *geojson.Feature) error {
	_ = "STUB: not implemented"
	return nil
}

func addSingleGeometryFeature(layer *vectortile.Tile_Layer, kve *keyValueEncoder, g orb.Geometry, p geojson.Properties, id any) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeProperties(kve *keyValueEncoder, properties geojson.Properties) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertID(id any) *uint64 { _ = "STUB: not implemented"; return nil }

func convertIntID(i int) *uint64 { _ = "STUB: not implemented"; return nil }
