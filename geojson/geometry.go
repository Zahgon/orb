package geojson

import (
	"errors"

	"github.com/paulmach/orb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ErrInvalidGeometry will be returned if the json of the geometry is invalid.
var ErrInvalidGeometry = errors.New("geojson: invalid geometry")

// A Geometry matches the structure of a GeoJSON Geometry.
type Geometry struct {
	Type        string       `json:"type"`
	Coordinates orb.Geometry `json:"coordinates,omitempty"`
	Geometries  []*Geometry  `json:"geometries,omitempty"`
}

// NewGeometry will create a Geometry object but will convert
// the input into a GeoJSON geometry. For example, it will convert
// Rings and Bounds into Polygons.
func NewGeometry(g orb.Geometry) *Geometry { _ = "STUB: not implemented"; return nil }

// Geometry returns the orb.Geometry for the geojson Geometry.
// This will convert the "Geometries" into an orb.Collection if applicable.
func (g *Geometry) Geometry() orb.Geometry { _ = "STUB: not implemented"; return *new(orb.Geometry) }

// MarshalJSON will marshal the geometry into the correct JSON structure.
func (g *Geometry) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the geometry into a BSON document with the structure
// of a GeoJSON Geometry. This function is used when the geometry is the top level
// document to be marshalled.
func (g *Geometry) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSONValue will marshal the geometry into a BSON value
// with the structure of a GeoJSON Geometry.
func (g *Geometry) MarshalBSONValue() (byte, []byte, error) {
	_ = "STUB: not implemented"
	// implementing MarshalBSONValue allows us to marshal into a null value
	// needed to match behavior with the JSON marshalling.
	return 0, nil, nil
}

func newGeometryMarshallDoc(g *Geometry) *geometryMarshallDoc {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalGeometry decodes the JSON data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(g) directly for the same result.
func UnmarshalGeometry(data []byte) (*Geometry, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the correct geometry from the JSON structure.
func (g *Geometry) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal a BSON document created with bson.Marshal.
func (g *Geometry) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A Point is a helper type that will marshal to/from a GeoJSON Point geometry.
type Point orb.Point

// Geometry will return the orb.Geometry version of the data.
func (p Point) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON will convert the Point into a GeoJSON Point geometry.
	new(orb.Geometry)
}

func (p Point) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the Point into a BSON value following the GeoJSON Point structure.
func (p Point) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the GeoJSON Point geometry.
func (p *Point) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal GeoJSON Point geometry.
func (p *Point) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A MultiPoint is a helper type that will marshal to/from a GeoJSON MultiPoint geometry.
type MultiPoint orb.MultiPoint

// Geometry will return the orb.Geometry version of the data.
func (mp MultiPoint) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON will convert the MultiPoint into a GeoJSON MultiPoint geometry.
	new(orb.Geometry)
}

func (mp MultiPoint) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the MultiPoint into a GeoJSON MultiPoint geometry BSON.
func (mp MultiPoint) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mp *MultiPoint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal the GeoJSON MultiPoint geometry.
func (mp *MultiPoint) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A LineString is a helper type that will marshal to/from a GeoJSON LineString geometry.
type LineString orb.LineString

// Geometry will return the orb.Geometry version of the data.
func (ls LineString) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON will convert the LineString into a GeoJSON LineString geometry.
	new(orb.Geometry)
}

func (ls LineString) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the LineString into a GeoJSON LineString geometry.
func (ls LineString) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (ls *LineString) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal the GeoJSON MultiPoint geometry.
func (ls *LineString) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A MultiLineString is a helper type that will marshal to/from a GeoJSON MultiLineString geometry.
type MultiLineString orb.MultiLineString

// Geometry will return the orb.Geometry version of the data.
func (mls MultiLineString) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

// MarshalJSON will convert the MultiLineString into a GeoJSON MultiLineString geometry.
func (mls MultiLineString) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalBSON will convert the MultiLineString into a GeoJSON MultiLineString geometry.
func (mls MultiLineString) MarshalBSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mls *MultiLineString) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal the GeoJSON MultiPoint geometry.
func (mls *MultiLineString) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A Polygon is a helper type that will marshal to/from a GeoJSON Polygon geometry.
type Polygon orb.Polygon

// Geometry will return the orb.Geometry version of the data.
func (p Polygon) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON will convert the Polygon into a GeoJSON Polygon geometry.
	new(orb.Geometry)
}

func (p Polygon) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the Polygon into a GeoJSON Polygon geometry.
func (p Polygon) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the GeoJSON Polygon geometry.
func (p *Polygon) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal the GeoJSON Polygon geometry.
func (p *Polygon) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// A MultiPolygon is a helper type that will marshal to/from a GeoJSON MultiPolygon geometry.
type MultiPolygon orb.MultiPolygon

// Geometry will return the orb.Geometry version of the data.
func (mp MultiPolygon) Geometry() orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

// MarshalJSON will convert the MultiPolygon into a GeoJSON MultiPolygon geometry.
func (mp MultiPolygon) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON will convert the MultiPolygon into a GeoJSON MultiPolygon geometry.
func (mp MultiPolygon) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON will unmarshal the GeoJSON MultiPolygon geometry.
func (mp *MultiPolygon) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal the GeoJSON MultiPolygon geometry.
func (mp *MultiPolygon) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type bsonGeometry struct {
	Type        string        `json:"type" bson:"type"`
	Coordinates bson.RawValue `json:"coordinates" bson:"coordinates"`
	Geometries  []*Geometry   `json:"geometries,omitempty" bson:"geometries"`
}

type jsonGeometry struct {
	Type        string           `json:"type"`
	Coordinates nocopyRawMessage `json:"coordinates"`
	Geometries  []*Geometry      `json:"geometries,omitempty"`
}

type geometryMarshallDoc struct {
	Type        string       `json:"type" bson:"type"`
	Coordinates orb.Geometry `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	Geometries  []*Geometry  `json:"geometries,omitempty" bson:"geometries,omitempty"`
}
