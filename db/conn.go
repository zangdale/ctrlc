package db

import (
	"context"
	"ctrlc/ent"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var client *ent.Client

var dbFile = "file:ent.db?_fk=1"

func init() {
	f := os.Getenv("ctrlc_db_file")
	if f != "" {
		dbFile = f
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("pwd")
	}

	dbFile = fmt.Sprintf("file:%s?_fk=1", filepath.Join(dir, "ent.db"))
}

func Init() {
	var err error
	client, err = ent.Open("sqlite3", dbFile)
	if err != nil {
		log.Panic().Err(err).Msg("failed opening connection to sqlite")
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}
}

func Close() error {
	if client == nil {
		return nil
	}
	return client.Close()
}
