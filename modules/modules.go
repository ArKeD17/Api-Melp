package modules

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// PageInfo Información de las paginas
type PageInfo struct {
	BackPage   int `json:"back_page"`
	NextPage   int `json:"next_page"`
	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
}

// PageInfoType Información de las paginas
var PageInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PageInfo",
	Description: "Información de las paginas",
	Fields: graphql.Fields{
		"back_page": &graphql.Field{
			Type:        graphql.Int,
			Description: "Numero de pagana anterior",
		},
		"next_page": &graphql.Field{
			Type:        graphql.Int,
			Description: "Numero de pagana siguiente",
		},
		"page": &graphql.Field{
			Type:        graphql.Int,
			Description: "Numero de pagana actual",
		},
		"total_pages": &graphql.Field{
			Type:        graphql.Int,
			Description: "Numero de total de paginas.",
		},
	},
})

// TokenScalar El `Token` representa una cadena que tiene la función de validar la sesión  del usuario, el token se enviá en los encabezados de la petición con el nombre de **Authorization**
type TokenScalar struct {
	Value string
}

func (scalar *TokenScalar) String() string {
	return scalar.Value
}

// NewTokenScalar El `Token` representa una cadena que tiene la función de validar la sesión  del usuario, el token se enviá en los encabezados de la petición con el nombre de **Authorization**
func NewTokenScalar(v string) *TokenScalar {
	return &TokenScalar{Value: v}
}

// TokenScalarType El `Token` representa una cadena que tiene la función de validar la sesión  del usuario, el token se enviá en los encabezados de la petición con el nombre de **Authorization**
var TokenScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Token",
	Description: "El `Token` representa una cadena que tiene la función de validar la sesión  del usuario, el token se enviá en los encabezados de la petición con el nombre de **Authorization**",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case TokenScalar:
			return value.String()
		case *TokenScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case TokenScalar:
			return value.String()
		case *TokenScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return NewTokenScalar(valueAST.Value)
		default:
			return nil
		}
	},
})

// DateScalar Fecha el formato valido es: YYYY-MM-DD
type DateScalar struct {
	Value string
}

func (scalar *DateScalar) String() string {
	return strings.Split(scalar.Value, "T")[0]
}

// NewDateScalar Fecha el formato valido es: YYYY-MM-DD
func NewDateScalar(v string) *DateScalar {
	return &DateScalar{Value: v}
}

// DateScalarType Fecha el formato valido es: YYYY-MM-DD
var DateScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Date",
	Description: "Fecha el formato valido es: YYYY-MM-DD",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case DateScalar:
			return value.String()
		case *DateScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case DateScalar:
			return value.String()
		case *DateScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			re := regexp.MustCompile(`^\d{4}\-((02\-(0[1-9]|[12][0-8]|19))|(((0[469])|11)\-(0[1-9]|[12][0-9]|30))|(((0[13578])|1[02])\-(0[1-9]|[12][0-9]|3[01])))$`)

			if !re.MatchString(valueAST.Value) {
				return nil
			}
			return NewDateScalar(valueAST.Value)
		default:
			return nil
		}
	},
})

// PhoneScalar Numero de teléfono, Los formatos validos son: +001122334455, 1122334455, 11-22-33-44-55, 11 22 33 44 55, +00-11-22-33-44-55, +00 11 22 33 44 55
type PhoneScalar struct {
	Value string
}

func (scalar *PhoneScalar) String() string {
	return scalar.Value
}

// NewPhoneScalar Numero de teléfono, Los formatos validos son: +001122334455, 1122334455, 11-22-33-44-55, 11 22 33 44 55, +00-11-22-33-44-55, +00 11 22 33 44 55
func NewPhoneScalar(v string) *PhoneScalar {
	return &PhoneScalar{Value: v}
}

// PhoneScalarType Numero de teléfono, Los formatos validos son: +001122334455, 1122334455, 11-22-33-44-55, 11 22 33 44 55, +00-11-22-33-44-55, +00 11 22 33 44 55
var PhoneScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Phone",
	Description: "Numero de teléfono, Los formatos validos son: +001122334455, 1122334455, 11-22-33-44-55, 11 22 33 44 55, +00-11-22-33-44-55, +00 11 22 33 44 55",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case PhoneScalar:
			return value.String()
		case *PhoneScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case PhoneScalar:
			return value.String()
		case *PhoneScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			re := regexp.MustCompile(`^(\+\d{2})?[ -]?(\(\d{3}\))?[ -]?([ -]?\d{2}){5}$`)

			if !re.MatchString(valueAST.Value) {
				return nil
			}
			return NewPhoneScalar(valueAST.Value)
		default:
			return nil
		}
	},
})

// EmailScalar Correo electrónico el formato es valido: correo@electronic.o
type EmailScalar struct {
	Value string
}

func (scalar *EmailScalar) String() string {
	return scalar.Value
}

// NewEmailScalar Correo electrónico el formato es valido: correo@electronic.o
func NewEmailScalar(v string) *EmailScalar {
	return &EmailScalar{Value: v}
}

// EmailScalarType Correo electrónico el formato es valido: correo@electronic.o
var EmailScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Email",
	Description: "Correo electrónico el formato es valido: correo@electronic.o",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case EmailScalar:
			return value.String()
		case *EmailScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case EmailScalar:
			return value.String()
		case *EmailScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			if !re.MatchString(valueAST.Value) {
				return nil
			}
			return NewEmailScalar(valueAST.Value)
		default:
			return nil
		}
	},
})

// PasswordScalar Contraseña para usuario debe tener una longitud minima de 8 caracteres ademas de contener al menos con un numero y una letra mayúscula
type PasswordScalar struct {
	Value string
}

func (scalar *PasswordScalar) String() string {
	return scalar.Value
}

// NewPasswordScalar Contraseña para usuario debe tener una longitud minima de 8 caracteres ademas de contener al menos con un numero y una letra mayúscula
func NewPasswordScalar(v string) *PasswordScalar {
	return &PasswordScalar{Value: v}
}

// PasswordScalarType Contraseña para usuario debe tener una longitud minima de 8 caracteres ademas de contener al menos con un numero y una letra mayúscula
var PasswordScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Password",
	Description: "Contraseña para usuario debe tener una longitud minima de 8 caracteres ademas de contener al menos con un numero y una letra mayúscula",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case PasswordScalar:
			return value.String()
		case *PasswordScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case PasswordScalar:
			return value.String()
		case *PasswordScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			number := false
			upper := false
			for _, c := range valueAST.Value {
				switch {
				case unicode.IsNumber(c):
					number = true
					break
				case unicode.IsUpper(c):
					upper = true
					break
				}
			}

			if !number || !upper || len(valueAST.Value) < 8 {
				return nil
			}
			return NewPasswordScalar(valueAST.Value)
		default:
			return nil
		}
	},
})

// TimeScalar Hora el formato valido es: hh:mm:ss
type TimeScalar struct {
	Value string
}

func (scalar *TimeScalar) String() string {
	scalar.Value = strings.Replace(scalar.Value, "Z", "", -1)
	return strings.Replace(scalar.Value, strings.Split(scalar.Value, "T")[0]+"T", "", -1)
}

// NewTimeScalar Hora el formato valido es: hh:mm:ss
func NewTimeScalar(v string) *TimeScalar {
	return &TimeScalar{Value: v}
}

// TimeScalarType Hora el formato valido es: hh:mm:ss
var TimeScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Time",
	Description: "Hora el formato valido es: hh:mm:ss",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case TimeScalar:
			return value.String()
		case *TimeScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case TimeScalar:
			return value.String()
		case *TimeScalar:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			re := regexp.MustCompile(`^([01]\d|2[0-3])(:[0-5]\d){2}$`)

			if !re.MatchString(valueAST.Value) {
				return nil
			}
			return NewTimeScalar(valueAST.Value)
		default:
			return nil
		}
	},
})
