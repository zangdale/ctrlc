package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"modernc.org/sqlite"
)

var dbc *sqlx.DB
var dbFile = "file:ent.db?_fk=1"

func init() {
	sql.Register("sqlite3", &sqlite.Driver{})
	f := os.Getenv("ctrlc_db_file")
	if f != "" {
		dbFile = f
	}
}

func Init(homeDir string) {
	var err error
	dbFile = fmt.Sprintf("file:%s?_fk=1", filepath.Join(homeDir, "ctrlc.db"))
	log.Info().Str("dbFile", dbFile).Msg("show db")

	dbc, err = sqlx.Connect("sqlite3", dbFile)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to sqlite")
	}
	dbc.MustExec(schema)
}

func Close() {
	if dbc == nil {
		return
	}
	dbc.Close()
}
