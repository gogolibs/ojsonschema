package ojsonschema_test

import (
	"encoding/json"
	"github.com/gogolibs/ojsonschema"
	"github.com/stretchr/testify/require"
	"testing"
)

type testCase struct {
	expected string
	actual   interface{}
}

var testCases = []testCase{
	{
		expected: `{ "type": "string" }`,
		actual:   ojsonschema.String{},
	},
	{
		expected: `{ "type": "string", "format": "regex" }`,
		actual:   ojsonschema.String{Format: "regex"},
	},
	{
		expected: `{ "type": "object" }`,
		actual:   ojsonschema.Object{},
	},
	{
		expected: `{ "type": "object", "additionalProperties": false }`,
		actual: ojsonschema.Object{
			AdditionalProperties: false,
		},
	},
	{
		expected: `{ "type": "array", "items": { "type": "string" } }`,
		actual: ojsonschema.Array{
			Items: ojsonschema.String{},
		},
	},
	{
		expected: `{ "oneOf": [ { "type": "string" }, { "type": "object" }] }`,
		actual: ojsonschema.OneOf(
			ojsonschema.String{},
			ojsonschema.Object{},
		),
	},
}

func TestCases(t *testing.T) {
	for _, testCase := range testCases {
		expectedI := new(interface{})
		err := json.Unmarshal([]byte(testCase.expected), expectedI)
		require.NoError(t, err)
		expectedNormalizedData, err := json.Marshal(expectedI)
		require.NoError(t, err)
		actualData, err := json.Marshal(testCase.actual)
		require.NoError(t, err)
		require.Equal(t, string(expectedNormalizedData), string(actualData))
	}
}
