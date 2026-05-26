package geojson

// Properties defines the feature properties with some helper methods.
type Properties map[string]any

// MustBool guarantees the return of a `bool` (with optional default).
// This function useful when you explicitly want a `bool` in a single
// value return context, for example:
//
//	myFunc(f.Properties.MustBool("param1"), f.Properties.MustBool("optional_param", true))
//
// This function will panic if the value is present but not a bool.
func (p Properties) MustBool(key string, def ...bool) bool { _ = "STUB: not implemented"; return false }

// MustInt guarantees the return of an `int` (with optional default).
// This function useful when you explicitly want a `int` in a single
// value return context, for example:
//
//	myFunc(f.Properties.MustInt("param1"), f.Properties.MustInt("optional_param", 123))
//
// This function will panic if the value is present but not a number.
func (p Properties) MustInt(key string, def ...int) int { _ = "STUB: not implemented"; return 0 }

// MustFloat64 guarantees the return of a `float64` (with optional default)
// This function useful when you explicitly want a `float64` in a single
// value return context, for example:
//
//	myFunc(f.Properties.MustFloat64("param1"), f.Properties.MustFloat64("optional_param", 10.1))
//
// This function will panic if the value is present but not a number.
func (p Properties) MustFloat64(key string, def ...float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// MustString guarantees the return of a `string` (with optional default)
// This function useful when you explicitly want a `string` in a single
// value return context, for example:
//
//	myFunc(f.Properties.MustString("param1"), f.Properties.MustString("optional_param", "default"))
//
// This function will panic if the value is present but not a string.
func (p Properties) MustString(key string, def ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// Clone returns a shallow copy of the properties.
func (p Properties) Clone() Properties { _ = "STUB: not implemented"; return *new(Properties) }
