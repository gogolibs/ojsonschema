package ojsonschema

import (
	"encoding/json"
	"github.com/gogolibs/ojson"
)

// Object represents object jsonschema instance
// {"type": "object", ... }
// with all relevant properties
type Object struct {
	Properties           ojson.Anything
	Required             ojson.Anything
	AdditionalProperties ojson.Anything
}

// MarshalJSON converts Object to a corresponding jsonschema object
func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":                 "object",
		"additionalProperties": o.AdditionalProperties,
		"properties":           o.Properties,
		"required":             o.Required,
	}))
}

// Array represents array jsonschema instance
// {"type": "array", ... }
// with all relevant properties
type Array struct {
	Items ojson.Anything
}

// MarshalJSON converts Array to a corresponding jsonschema object
func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":  "array",
		"items": a.Items,
	}))
}

// String represents string jsonschema instance
// {"type": "string", ... }
// with all relevant properties
type String struct {
	Enum   ojson.Anything
	Format ojson.Anything
}

// MarshalJSON converts String to a corresponding jsonschema object
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":   "string",
		"enum":   s.Enum,
		"format": s.Format,
	}))
}

func removeNilValues(obj ojson.Object) ojson.Object {
	objWithoutNils := ojson.Object{}
	for key, value := range obj {
		if value != nil {
			objWithoutNils[key] = value
		}
	}
	return objWithoutNils
}

// Const represents const jsonschema instance with a single key
// {"const": <value>}
func Const(value ojson.Anything) ojson.Object {
	return ojson.Object{
		"const": value,
	}
}

// Enum represents const jsonschema instance with a single key
// {"enum": <value>}
func Enum(values ...ojson.Anything) ojson.Object {
	return ojson.Object{
		"enum": values,
	}
}

// OneOf returns JSON-marshallable {"oneOf": [<schemas>]} jsonschema object
func OneOf(schemas ...ojson.Anything) ojson.Object {
	return ojson.Object{
		"oneOf": schemas,
	}
}

// Integer represents integer jsonschema instance
// {"type": "integer", ... }
// with all relevant properties
type Integer struct {
	Enum ojson.Anything
}

// MarshalJSON converts Integer to a corresponding jsonschema object
func (i Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type": "integer",
		"enum": i.Enum,
	}))
}

// Number represents number jsonschema instance
// {"type": "number", ... }
// with all relevant properties
type Number struct {
	Enum ojson.Anything
}

// MarshalJSON converts Number to a corresponding jsonschema object
func (n Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type": "number",
		"enum": n.Enum,
	}))
}
