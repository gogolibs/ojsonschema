package ojsonschema_test

import (
	"github.com/gogolibs/ojson"
	"github.com/gogolibs/ojsonschema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExample(t *testing.T) {
	schema := ojsonschema.Object{
		AdditionalProperties: false,
		Properties: ojson.Object{
			"requiredField": ojsonschema.String{},
			"optionalField": ojsonschema.String{},
		},
		Required: ojson.Array{"requiredField"},
	}
	expected := string(ojson.MustMarshal(ojson.Object{
		"type":                 "object",
		"additionalProperties": false,
		"properties": ojson.Object{
			"requiredField": ojson.Object{"type": "string"},
			"optionalField": ojson.Object{"type": "string"},
		},
		"required": ojson.Array{"requiredField"},
	}))
	actual := string(ojson.MustMarshal(schema))
	require.Equal(t, expected, actual)
}
