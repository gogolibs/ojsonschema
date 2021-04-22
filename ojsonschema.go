package ojsonschema

import (
	"encoding/json"
	"github.com/gogolibs/ojson"
)

type Object struct {
	Properties           ojson.Anything
	Required             ojson.Anything
	AdditionalProperties ojson.Anything
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":                 "object",
		"additionalProperties": o.AdditionalProperties,
		"properties":           o.Properties,
		"required":             o.Required,
	}))
}

type Array struct {
	Items ojson.Anything
}

func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":  "array",
		"items": a.Items,
	}))
}

type String struct {
	Format ojson.Anything
}

func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(removeNilValues(ojson.Object{
		"type":   "string",
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

func OneOf(schemas ...ojson.Anything) ojson.Object {
	return ojson.Object{
		"oneOf": schemas,
	}
}
