package sumoapi

import (
	"maps"
	"reflect"

	"github.com/google/jsonschema-go/jsonschema"
)

var typeSchemas = make(map[reflect.Type]*jsonschema.Schema)

// TypeSchemas returns the registered type schemas.
func TypeSchemas() map[reflect.Type]*jsonschema.Schema {
	return maps.Clone(typeSchemas)
}
