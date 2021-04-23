package ojsonschema_test

import (
	"encoding/json"
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
		name: "string: simple",
		expected: `{ "type": "string" }`,
		actual:   ojsonschema.String{},
	},
	{
		name: "string: format regex",
		expected: `{ "type": "string", "format": "regex" }`,
		actual:   ojsonschema.String{Format: "regex"},
	},
	{
		name: "object: simple",
		expected: `{ "type": "object" }`,
		actual:   ojsonschema.Object{},
	},
	{
		name: "object: additional properties flag",
		expected: `{ "type": "object", "additionalProperties": false }`,
		actual: ojsonschema.Object{
			AdditionalProperties: false,
		},
	},
	{
		name: "array: strings",
		expected: `{ "type": "array", "items": { "type": "string" } }`,
		actual: ojsonschema.Array{
			Items: ojsonschema.String{},
		},
	},
	{
		name: "oneOf: string, object",
		expected: `{ "oneOf": [ { "type": "string" }, { "type": "object" }] }`,
		actual: ojsonschema.OneOf(
			ojsonschema.String{},
			ojsonschema.Object{},
		),
	},
}

func TestCases(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedI := new(interface{})
			err := json.Unmarshal([]byte(testCase.expected), expectedI)
			require.NoError(t, err)
			expectedNormalizedData, err := json.Marshal(expectedI)
			require.NoError(t, err)
			actualData, err := json.Marshal(testCase.actual)
			require.NoError(t, err)
			require.Equal(t, string(expectedNormalizedData), string(actualData))
		})
	}
}
