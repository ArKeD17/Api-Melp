package schema

import (
	"github.com/graphql-go/graphql"
	restaurant "gitlab.com/melp/api/modules/restaurant"
)

// RootMutation comment
var rootMutation graphql.ObjectConfig

func init() {
	rootMutation = graphql.ObjectConfig{
		Name:        "RootMutation",
		Description: "Mutaciones con las que cuenta la API",
		Fields:      mergeFields(restaurant.Mutations),
	}
}

// GetMutationObject comment
func GetMutationObject() *graphql.Object {
	return graphql.NewObject(rootMutation)
}
