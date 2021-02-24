package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
)

// Schema comment
var Schema graphql.Schema

func init() {
	var err error

	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    GetQueryObject(),
		Mutation: GetMutationObject(),
	})
	logger.Error("NEW SCHEMA", err)
}

func mergeFields(newFields ...graphql.Fields) graphql.Fields {
	fields := graphql.Fields{}
	for _, field := range newFields {
		for k, v := range field {
			fields[k] = v
		}
	}

	return fields
}
