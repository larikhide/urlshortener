package pgstore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/larikhide/urlshortener/app/repos/urls"
)

var _ urls.URLStore = &PostgresDB{}

type DBPgUrl struct {
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	LongURL   string     `db:"short_url"`
	ShortURL  string     `db:"long_url"`
}

type PostgresDB struct {
	db *sql.DB
}

func NewUsers(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS public.urls (
		created_at timestamptz NOT NULL,
		updated_at timestamptz NOT NULL,
		deleted_at timestamptz NULL,
		short_url varchar NOT NULL,
		"long_url" varchar NULL,
		CONSTRAINT users_pk PRIMARY KEY (short_url)
	);`)

	if err != nil {
		db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	us := &PostgresDB{
		db: db,
	}
	return us, nil
}

func (us *PostgresDB) Close() {
	us.db.Close()
}

func (pgs *PostgresDB) SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string) error {
	dbu := &DBPgUrl{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ShortURL:  shortUrl,
		LongURL:   longUrl,
	}

	_, err := pgs.db.ExecContext(ctx, `INSERT INTO urls 
	(created_at, updated_at, deleted_at, short_url, long_url)
	values ($1, $2, $3, $4, $5)`,

		dbu.CreatedAt,
		dbu.UpdatedAt,
		nil,
		dbu.ShortURL,
		dbu.LongURL,
	)
	if err != nil {
		return fmt.Errorf("failing saving to postgres error: %v", err)
	}

	return nil
}

func (pgs *PostgresDB) RetrieveInitialUrl(ctx context.Context, shortUrl string) (*urls.URL, error) {
	dbu := &DBPgUrl{}
	rows, err := pgs.db.QueryContext(ctx, `SELECT created_at, updated_at, deleted_at, short_url, long_url 
	FROM users WHERE short_url = $1`, shortUrl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&dbu.CreatedAt,
			&dbu.UpdatedAt,
			&dbu.DeletedAt,
			&dbu.ShortURL,
			&dbu.LongURL,
		); err != nil {
			return nil, err
		}
	}

	return &urls.URL{
		LongURL:  dbu.LongURL,
		ShortURL: dbu.ShortURL,
	}, nil
}
