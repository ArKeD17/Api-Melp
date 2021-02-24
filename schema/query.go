package schema

import (
	"github.com/graphql-go/graphql"
	restaurant "gitlab.com/melp/api/modules/restaurant"
)

// RootQuery comment
var rootQuery graphql.ObjectConfig

func init() {
	rootQuery = graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Queries con las que cuenta la API",
		Fields: mergeFields(
			restaurant.Queries,
		),
	}
}

// GetQueryObject comment
func GetQueryObject() *graphql.Object {
	return graphql.NewObject(rootQuery)
}
