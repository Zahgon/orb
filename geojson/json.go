package geojson

// CustomJSONMarshaler can be set to have the code use a different
// json marshaler than the default in the standard library.
// One use case in enabling `github.com/json-iterator/go`
// with something like this:
//
//	import (
//	  jsoniter "github.com/json-iterator/go"
//	  "github.com/paulmach/orb/geojson"
//	)
//
//	// in an init() or main(), etc.
//	var c = jsoniter.Config{
//	  EscapeHTML:              true,
//	  SortMapKeys:             false,
//	  MarshalFloatWith6Digits: true,
//	}.Froze()
//
//	geojson.CustomJSONMarshaler = c
//	geojson.CustomJSONUnmarshaler = c
//
// Note that any errors encountered during marshaling will be different.
var CustomJSONMarshaler interface {
	Marshal(v any) ([]byte, error)
} = nil

// CustomJSONUnmarshaler can be set to have the code use a different
// json unmarshaler than the default in the standard library.
// One use case in enabling `github.com/json-iterator/go`
// with something like this:
//
//	import (
//	  jsoniter "github.com/json-iterator/go"
//	  "github.com/paulmach/orb/geojson"
//	)
//
//	// in an init() or main(), etc.
//	c := jsoniter.Config{
//	  EscapeHTML:              true,
//	  SortMapKeys:             false,
//	  MarshalFloatWith6Digits: true,
//	}.Froze()
//
//	geojson.CustomJSONMarshaler = c
//	geojson.CustomJSONUnmarshaler = c
//
// Note that any errors encountered during unmarshaling will be different.
var CustomJSONUnmarshaler interface {
	Unmarshal(data []byte, v any) error
} = nil

func marshalJSON(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func unmarshalJSON(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

type nocopyRawMessage []byte

func (m *nocopyRawMessage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
