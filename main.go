package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type RootResolver struct{}

func (r *RootResolver) Info() (string, error) {
	return "hahaha", errors.New("some error")
}

func parseSchema(path string, resolver interface{}) *graphql.Schema {
	astr, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	schemaString := string(astr)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
	)
	if err != nil {
		panic(err)
	}
	return parsedSchema
}

func main() {
	http.Handle("/graphql", &relay.Handler{
		Schema: parseSchema("./schema.graphql", &RootResolver{}),
	})

	fmt.Println("tangina mo")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
