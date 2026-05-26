package mvt

import (
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
)

const (
	// DefaultExtent for mapbox vector tiles. (https://www.mapbox.com/vector-tiles/specification/)
	DefaultExtent = 4096
)

// Layer is intermediate MVT layer to be encoded/decoded or projected.
type Layer struct {
	Name     string
	Version  uint32
	Extent   uint32
	Features []*geojson.Feature
}

// NewLayer is a helper to create a Layer from a feature collection
// and a name, it sets the default extent and version to 1.
func NewLayer(name string, fc *geojson.FeatureCollection) *Layer {
	_ = "STUB: not implemented"
	return nil
}

// ProjectToTile will project all the geometries in the layer
// to tile coordinates based on the extent and the mercator projection.
func (l *Layer) ProjectToTile(tile maptile.Tile) { _ = "STUB: not implemented"; return }

// ProjectToWGS84 will project all the geometries backed to WGS84 from
// the extent and mercator projection.
func (l *Layer) ProjectToWGS84(tile maptile.Tile) { _ = "STUB: not implemented"; return }

// Layers is a set of layers.
type Layers []*Layer

// NewLayers creates a set of layers given a set of feature collections.
func NewLayers(layers map[string]*geojson.FeatureCollection) Layers {
	_ = "STUB: not implemented"
	return *new(Layers)
}

// ToFeatureCollections converts the layers to sets of geojson
// feature collections.
func (ls Layers) ToFeatureCollections() map[string]*geojson.FeatureCollection {
	_ = "STUB: not implemented"
	return nil
}

// ProjectToTile will project all the geometries in all layers
// to tile coordinates based on the extent and the mercator projection.
func (ls Layers) ProjectToTile(tile maptile.Tile) { _ = "STUB: not implemented"; return }

// ProjectToWGS84 will project all the geometries in all the layers backed
// to WGS84 from the extent and mercator projection.
func (ls Layers) ProjectToWGS84(tile maptile.Tile) { _ = "STUB: not implemented"; return }
