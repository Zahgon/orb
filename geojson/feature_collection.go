/*
Package geojson is a library for encoding and decoding GeoJSON into Go structs
using the geometries in the orb package. Supports both the json.Marshaler and
json.Unmarshaler interfaces as well as helper functions such as
`UnmarshalFeatureCollection` and `UnmarshalFeature`.
*/
package geojson

const featureCollection = "FeatureCollection"

// A FeatureCollectionOf correlates to a GeoJSON feature collection but allows for a generic type
// for the properties of the features.
//
// The code assumes type of P is a struct, or map as the GeoJSON spec requires it
// marshal into the a json object.
type FeatureCollectionOf[P any] struct {
	Type     string          `json:"type"`
	BBox     BBox            `json:"bbox,omitempty"`
	Features []*FeatureOf[P] `json:"features"`

	// ExtraMembers can be used to encoded/decode extra key/members in
	// the base of the feature collection. Note that keys of "type", "bbox"
	// and "features" will not work as those are reserved by the GeoJSON spec.
	ExtraMembers Properties `json:"-"`
}

// A FeatureCollection correlates to a GeoJSON feature collection.
type FeatureCollection = FeatureCollectionOf[Properties]

// NewFeatureCollection creates and initializes a new feature collection.
func NewFeatureCollection() *FeatureCollection { _ = "STUB: not implemented"; return nil }

// Append appends a feature to the collection.
func (fc *FeatureCollectionOf[P]) Append(feature *FeatureOf[P]) *FeatureCollectionOf[P] {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON converts the feature collection object into the proper JSON.
// It will handle the encoding of all the child features and geometries.
// Alternately one can call json.Marshal(fc) directly for the same result.
// Items in the ExtraMembers map will be included in the base of the
// feature collection object.
func (fc FeatureCollectionOf[P]) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalBSON converts the feature collection object into a BSON document
// represented by bytes. It will handle the encoding of all the child features
// and geometries.
// Items in the ExtraMembers map will be included in the base of the
// feature collection object.
func (fc FeatureCollectionOf[P]) MarshalBSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFeatureCollectionDoc[P any](fc FeatureCollectionOf[P]) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON decodes the data into a GeoJSON feature collection.
// Extra/foreign members will be put into the `ExtraMembers` attribute.
func (fc *FeatureCollectionOf[P]) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalBSON will unmarshal a BSON document created with bson.Marshal.
// Extra/foreign members will be put into the `ExtraMembers` attribute.
func (fc *FeatureCollectionOf[P]) UnmarshalBSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalFeatureCollection decodes the data into a GeoJSON feature collection.
// Alternately one can call json.Unmarshal(fc) directly for the same result.
func UnmarshalFeatureCollection(data []byte) (*FeatureCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
