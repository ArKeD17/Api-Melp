package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/graphql-go/graphql"
	"gitlab.com/melp/api/schema"
)

// Run comment
func Run() {
	port := ":" + os.Getenv("PORT")

	http.HandleFunc("/graphql", CORS(httpGraphql))
	log.Println(os.Getenv("PORT"))
	log.Printf("Servidor escuchando en: http://localhost%s", port)
	http.ListenAndServe(port, nil)
}

func httpGraphql(w http.ResponseWriter, r *http.Request) {
	var query string

	switch strings.ToUpper(r.Method) {
	case "GET":
		query = r.URL.Query().Get("query")
		break
	default:
		switch r.Header.Get("Content-Type") {
		case "application/json":
			defer r.Body.Close()
			bytes, _ := ioutil.ReadAll(r.Body)

			var body map[string]interface{}
			json.Unmarshal(bytes, &body)
			query = body["query"].(string)
		default:
			query = r.FormValue("query")
		}
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// CORS Access control allow origin
func CORS(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Accept, Origin, Authorization")
			return
		}
		h(w, r)
	})
}
