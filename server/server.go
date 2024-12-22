package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	ientgql "github.com/naoto67/entgql"
	"github.com/naoto67/entgql/ent"
)

func main() {
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	client = client.Debug()
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("running schema migration", err)
	}

	srv := handler.NewDefaultServer(ientgql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)

	server := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("http server terminated", err)
	}
}
