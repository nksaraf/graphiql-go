# graphiql-go

A Http handler for your GraphQL servers in Go that renders the GraphiQL playground for you to test your server.

### Installation
```
go get github.com/nksaraf/graphiql-go
```

## Example
```go
package main

import (
        "log"
        "net/http"

        graphql "github.com/graph-gophers/graphql-go"
        "github.com/graph-gophers/graphql-go/relay"
        "github.com/nksaraf/graphiql-go"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func main() {
        s := `
                type Query {
                        hello: String!
                }
        `
        schema := graphql.MustParseSchema(s, &query{})
        http.Handle("/query", &relay.Handler{Schema: schema})
        
        // Assign the graphiql handler to your endpoint of choice
        http.Handle("/graphiql", &graphiql.Handler{Url:"/query"})
        
        http.ListenAndServe(":8080", nil)
}
```

Now, open your browser and navigate to `http://localhost:8080/graphiql`. 
Click the *Explorer* button to open the GraphiQL Explorer and enjoy the amazing experience of building your query while you explore your graph.
