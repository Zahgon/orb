package geojson

import (
	"encoding/json"

	"github.com/paulmach/orb"
)

// A FeatureOf corresponds to GeoJSON feature object but allows for a generic type for the properties.
// This allows users to unmarshal into a struct instead of a map if they choose.
//
// The code assumes type of P is a struct, or map as the GeoJSON spec requires it
// marshal into the a json object.
type FeatureOf[P any] struct {
	ID         any          `json:"id,omitempty"`
	Type       string       `json:"type"`
	BBox       BBox         `json:"bbox,omitempty"`
	Geometry   orb.Geometry `json:"geometry"`
	Properties P            `json:"properties"`

	// ExtraMembers can be used to encoded/decode extra key/members in
	// the base of the feature object. Note that keys of "id", "type", "bbox"
	// "geometry" and "properties" will not work as those are reserved by the
	// GeoJSON spec.
	ExtraMembers Properties `json:"-"`
}

// A Feature corresponds to GeoJSON feature object.
type Feature = FeatureOf[Properties]

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry orb.Geometry) *Feature { _ = "STUB: not implemented"; return nil }

// Point implements the orb.Pointer interface so that Features can be used
// with quadtrees. The point returned is the center of the Bound of the geometry.
// To represent the geometry with another point you must create a wrapper type.
func (f *FeatureOf[P]) Point() orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

var _ orb.Pointer = &FeatureOf[any]{}

// MarshalJSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
// Items in the ExtraMembers map will be included in the base of the
// feature object.
func (f FeatureOf[P]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
// Items in the ExtraMembers map will be included in the base of the
// feature object.
func (f FeatureOf[P]) MarshalBSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f FeatureOf[P]) jsonProperties() (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

// empty
// we assume it's an object so an empty {} is 2 bytes
// in that case the properties should be nil according to the geojson spec

func (f FeatureOf[P]) bsonProperties() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f FeatureOf[P]) newFeatureDoc(properties any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// UnmarshalFeature decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(f) directly for the same result.
func UnmarshalFeature(data []byte) (*Feature, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON handles the correct unmarshalling of the data
// into the orb.Geometry types.
func (f *FeatureOf[P]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalBSON will unmarshal a BSON document created with bson.Marshal.
func (f *FeatureOf[P]) UnmarshalBSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type featureDoc[P any] struct {
	ID         any       `json:"id,omitempty" bson:"id"`
	Type       string    `json:"type" bson:"type"`
	BBox       BBox      `json:"bbox,omitempty" bson:"bbox,omitempty"`
	Geometry   *Geometry `json:"geometry" bson:"geometry"`
	Properties P         `json:"properties" bson:"properties"`
}
