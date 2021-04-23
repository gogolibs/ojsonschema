# ojsonschema

[![GoDoc](https://godoc.org/github.com/gogolibs/ojsonschema?status.svg)](https://pkg.go.dev/github.com/gogolibs/ojsonschema)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogolibs/ojsonschema)](https://goreportcard.com/report/github.com/gogolibs/ojsonschema)
[![CI](https://github.com/gogolibs/ojsonschema/actions/workflows/test-and-coverage.yml/badge.svg)](https://github.com/gogolibs/ojsonschema/actions/workflows/test-and-coverage.yml)
[![Tests CI](https://github.com/gogolibs/ojsonschema-tests/actions/workflows/ci.yml/badge.svg)](https://github.com/gogolibs/ojsonschema-tests/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/gogolibs/ojsonschema/branch/main/graph/badge.svg?token=JXSDP6Ifxi)](https://codecov.io/gh/gogolibs/ojsonschema)

**ojsonschema** is a collection of convenience type aliases to represent JSON schema objects in plain Go.

**NOTE**: consider also using [ojson](https://github.com/gogolibs/ojson) library to represent JSON objects in Go.

**NOTE**: this library is used to only represent JSON schema objects, not to validate them. For the actual
implementation of json schema consider using one of the libraries developed specifically for that (for example,
[qri-io/jsonschema](https://github.com/qri-io/jsonschema) 
or [santhosh-tekuri/jsonschema](https://github.com/santhosh-tekuri/jsonschema))

## Example usage: ##

```go
package example

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
```