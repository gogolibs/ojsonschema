package ojsonschema_test

import (
	"encoding/json"
	"github.com/gogolibs/ojson"
	"github.com/gogolibs/ojsonschema"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCases = []struct {
	name     string
	expected string
	actual   interface{}
}{
	{
		name:     "string: simple",
		expected: `{ "type": "string" }`,
		actual:   ojsonschema.String{},
	},
	{
		name:     "string: format regex",
		expected: `{ "type": "string", "format": "regex" }`,
		actual:   ojsonschema.String{Format: "regex"},
	},
	{
		name:     "string: enum",
		expected: `{ "type": "string", "enum": ["one", "two", "three"] }`,
		actual:   ojsonschema.String{Enum: ojson.Array{"one", "two", "three"}},
	},
	{
		name:     "integer: simple",
		expected: `{ "type": "integer" }`,
		actual:   ojsonschema.Integer{},
	},
	{
		name:     "integer: enum",
		expected: `{ "type": "integer", "enum": [1, 2, 3] }`,
		actual:   ojsonschema.Integer{Enum: ojson.Array{1, 2, 3}},
	},
	{
		name:     "number: simple",
		expected: `{ "type": "number" }`,
		actual:   ojsonschema.Number{},
	},
	{
		name:     "number: enum",
		expected: `{ "type": "number", "enum": [1, 2.0, 3] }`,
		actual:   ojsonschema.Number{Enum: ojson.Array{1, 2.0, 3}},
	},
	{
		name:     "object: simple",
		expected: `{ "type": "object" }`,
		actual:   ojsonschema.Object{},
	},
	{
		name:     "object: additional properties flag",
		expected: `{ "type": "object", "additionalProperties": false }`,
		actual: ojsonschema.Object{
			AdditionalProperties: false,
		},
	},
	{
		name:     "array: strings",
		expected: `{ "type": "array", "items": { "type": "string" } }`,
		actual: ojsonschema.Array{
			Items: ojsonschema.String{},
		},
	},
	{
		name:     "oneOf: string, object",
		expected: `{ "oneOf": [ { "type": "string" }, { "type": "object" }] }`,
		actual: ojsonschema.OneOf(
			ojsonschema.String{},
			ojsonschema.Object{},
		),
	},
	{
		name:     "enum",
		expected: `{ "enum": [ "one", 1, null ] }`,
		actual:   ojsonschema.Enum("one", 1, nil),
	},
	{
		name:     "const",
		expected: `{ "const": "hello" }`,
		actual:   ojsonschema.Const("hello"),
	},
	{
		name:     "ref",
		expected: `{ "$ref": "#/definitions/some" }`,
		actual:   ojsonschema.Ref("#/definitions/some"),
	},
}

func TestCases(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedI := new(interface{})
			err := json.Unmarshal([]byte(testCase.expected), expectedI)
			require.NoError(t, err)
			expectedNormalizedData := ojson.MustMarshal(expectedI)
			actualData := ojson.MustMarshal(testCase.actual)
			require.Equal(t, string(expectedNormalizedData), string(actualData))
		})
	}
}
