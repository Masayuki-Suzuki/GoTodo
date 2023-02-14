package database

import (
  "app/ent"
  "app/ent/migrate"
  "context"
  "log"
)

const (
  dsn = "admin:admin@tcp(mariadb:3306)/go_todo?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
  Client *ent.Client
)

func Connect() {
  db, err := ent.Open("mysql", dsn)

  if err != nil {
    log.Fatalf("Database connection error occurred: %v\n", err)
  }

  Client = db
  ctx := context.Background()

  // Run the auto migration tool.
  if err := Client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
    log.Fatalf("Failed creating schema resources: %v", err)
  }
}
