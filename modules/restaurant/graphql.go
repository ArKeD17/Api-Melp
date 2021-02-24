package restaurant

import (
	"github.com/graphql-go/graphql"
	"gitlab.com/melp/api/modules"
)

// Restaurant Tipo de dato para los restaurantes
type Restaurant struct {
	ID     string               `json:"id"`
	Rating int                  `json:"rating"`
	Name   string               `json:"name"`
	Site   string               `json:"site"`
	Email  *modules.EmailScalar `json:"email"`
	Phone  *modules.PhoneScalar `json:"phone"`
	Street string               `json:"street"`
	City   string               `json:"city"`
	State  string               `json:"state"`
	Lat    string               `json:"lat"`
	Lng    string               `json:"lng"`
}

// StatisticsRestaurant Tipo de dato para las estadisticas de los restaurantes
type StatisticsRestaurant struct {
	Count int     `json:"count"`
	Avg   float64 `json:"avg"`
	Std   float64 `json:"std"`
}

// ListRestaurants Lista de restaurantes
type ListRestaurants struct {
	Data []Restaurant `json:"data"`
}

// Queries comment
var Queries = graphql.Fields{
	"getRestaurant": &graphql.Field{
		Type:        graphql.NewNonNull(RestaurantType),
		Description: "Retorna los datos de un restaurante en especifico.",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: Restaurant{}.GetRestaurant,
	},
	"restaurants": &graphql.Field{
		Type:        graphql.NewNonNull(ListRestaurantsType),
		Description: "Retorna la lista de restaurantes",
		Resolve:     Restaurant{}.Restaurants,
	},
	"statisticsRestaurants": &graphql.Field{
		Type:        graphql.NewNonNull(StatisticsRestaurantType),
		Description: "Retorna estadisticas de ciertos restaurantes basado en un radio medido en metros.",
		Args: graphql.FieldConfigArgument{
			"lat": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lng": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"radius": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: StatisticsRestaurant{}.StatisticsRestaurants,
	},
}

// Mutations comment
var Mutations = graphql.Fields{
	"createRestaurant": &graphql.Field{
		Type:        graphql.NewNonNull(RestaurantType),
		Description: "Registra un nuevo restaurante",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"rating": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"site": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(modules.EmailScalarType),
			},
			"phone": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(modules.PhoneScalarType),
			},
			"street": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"city": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"state": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lat": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lng": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: Restaurant{}.Create,
	},
	"updateRestaurant": &graphql.Field{
		Type:        graphql.NewNonNull(RestaurantType),
		Description: "Actualiza los datos de un restaurante en especifico",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"rating": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"site": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"email": &graphql.ArgumentConfig{
				Type: modules.EmailScalarType,
			},
			"phone": &graphql.ArgumentConfig{
				Type: modules.PhoneScalarType,
			},
			"street": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"city": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"state": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lat": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lng": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: Restaurant{}.Update,
	},
	"deleteRestaurant": &graphql.Field{
		Type:        graphql.NewNonNull(graphql.String),
		Description: "Elimina el registro de un restaurante en especifico",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: Restaurant{}.Delete,
	},
}

// RestaurantType Tipo de dato para los restaurantes
var RestaurantType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Restaurant",
	Description: "Tipo de dato para los restaurantes",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Identificador del restaurante",
		},
		"rating": &graphql.Field{
			Type:        graphql.Int,
			Description: "Calificación del restaurante",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"site": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"email": &graphql.Field{
			Type:        modules.EmailScalarType,
			Description: "Correo electrónico del cliente",
		},
		"phone": &graphql.Field{
			Type:        modules.PhoneScalarType,
			Description: "Numero telefónico del restaurante.",
		},
		"street": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"city": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"state": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"lat": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
		"lng": &graphql.Field{
			Type:        graphql.String,
			Description: "Nombre del restaurante",
		},
	},
})

// StatisticsRestaurantType Tipo de dato para las estadisticas de los restaurantes
var StatisticsRestaurantType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "StatisticsRestaurant",
	Description: "Tipo de dato para las estadisticas de los restaurantes",
	Fields: graphql.Fields{
		"count": &graphql.Field{
			Type:        graphql.Int,
			Description: "Cantidad de restaurante que se encuentran dentro del radio.",
		},
		"avg": &graphql.Field{
			Type:        graphql.Float,
			Description: "Promedio de Calificación de los restaurante",
		},
		"std": &graphql.Field{
			Type:        graphql.Float,
			Description: "Desviación estandar de la calificación de los restaurantes dentro del radio.",
		},
	},
})

// ListRestaurantsType Lista de restaurantes
var ListRestaurantsType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ListRestaurants",
	Description: "Lista de restaurantes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type:        graphql.NewList(graphql.NewNonNull(RestaurantType)),
			Description: "Lista de restaurantes",
		},
	},
})
