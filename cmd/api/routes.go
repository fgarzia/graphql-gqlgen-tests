package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/graphql-tests/internal/graphqls"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func mapRoutes(router *gin.Engine) {
	router.POST("/graphql", graphqlHandler())

	router.GET("/playground", playgroundHandler())

	router.GET("/ping", ping)
}

func ping(ginCtx *gin.Context) {
	txn := nrgin.Transaction(ginCtx)
	txn.Ignore()

	ginCtx.String(http.StatusOK, "pong")
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graphqls.NewExecutableSchema(graphqls.Config{Resolvers: &graphqls.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
