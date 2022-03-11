package db

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

var name, _ = os.Hostname()
var fromSystem = fmt.Sprintf("%s-%s-%s", runtime.GOOS, runtime.GOARCH, name)

var schema = `
CREATE TABLE IF NOT EXISTS copy_msgs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	body varchar(255) NOT NULL,
	'from' varchar(255) NULL,
	created_at datetime NOT NULL
)`

type Backups struct {
	ID       uint64    `db:"id"`
	Body     string    `db:"body"`
	From     string    `db:"from"`
	CreateAt time.Time `db:"created_at"`
}

func SaveCopyMsg(ctx context.Context, text string) error {
	_, err := dbc.MustExec("INSERT INTO copy_msgs (body, 'from', created_at) VALUES ($1,$2,$3)",
		text, fromSystem, time.Now()).LastInsertId()
	return err
}

/*

CREATE TABLE IF NOT EXISTS backups(
	id varchar(255) UNIQUE NOT NULL,
	title varchar(255) NOT NULL,
	format varchar(255) NOT NULL,
	url varchar(255) NOT NULL,
	hash varchar(255) NOT NULL,
	book_id integer NOT NULL,
	book_slug varchar(255) NOT NULL,
	book_name varchar(255) NOT NULL,
	create_at datetime NOT NULL,
	update_at datetime NOT NULL,
	PRIMARY KEY(id)
)

func create(b *Backups) error {
	_, err := dbc.MustExec("INSERT INTO backups (id, title, format,url,hash,book_id,book_slug,book_name,create_at,update_at) VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10)",
		b.ID, b.Title, b.Format, b.URL, b.Hash, b.BookID, b.BookSlug, b.BookName, b.CreateAt, b.UpdateAt).LastInsertId()
	return err
}

func update(b *Backups) error {
	_, err := dbc.MustExec("UPDATE backups SET title = $1, format = $2,url = $3,hash = $4,book_id = $5,book_slug = $6,book_name = $7,create_at = $8,update_at = $9) WHERE id = $10",
		b.Title, b.Format, b.URL, b.Hash, b.BookID, b.BookSlug, b.BookName, b.CreateAt, b.UpdateAt, b.ID).RowsAffected()
	return err
}

func get(id string) (*Backups, error) {
	res := &Backups{}
	err := dbc.Get(res, "SELECT * FROM backups WHERE id=$1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return res, err
}
*/
