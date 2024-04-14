package backend

import (
  "database/sql"
  "log"
  _ "github.com/mattn/go-sqlite3"
)

func CreateDb() *sql.DB {
  db, err := sql.Open("sqlite3", "./db/data.db")

  if err != nil {
    log.Fatal(err)
  }

  return db
}
