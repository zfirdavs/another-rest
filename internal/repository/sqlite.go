package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zfirdavs/another-rest/internal/entity"
)

type sqliteRepo struct {
	dbname string
	db     *sql.DB
}

func NewSQLiteRepository(dbname string) *sqliteRepo {
	return &sqliteRepo{dbname: dbname}
}

func (s *sqliteRepo) Init() error {
	// open the connection
	db, err := sql.Open("sqlite3", "./"+s.dbname)
	if err != nil {
		return fmt.Errorf("failed to initialize sqlite database: %w", err)
	}

	s.db = db
	return nil
}

func (s *sqliteRepo) CreatePostsTable() error {
	sqlStmt := `
		create table if not exists posts (
			id integer not null primary key,
			title text,
			txt text
		);
	`

	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("failed to create the posts table: %w", err)
	}

	// s.Close()
	return nil
}

func (s *sqliteRepo) Close() error {
	return s.db.Close()
}

func (s *sqliteRepo) Save(ctx context.Context, post *entity.Post) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin sqlite transaction: %w", err)
	}

	query := `insert into posts (id, title, txt) values (?, ?, ?)`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare sqlite transaction: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		return fmt.Errorf("failed to exec sqlite transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit sqlite transaction: %w", err)
	}
	return nil
}

func (s *sqliteRepo) FindAll(ctx context.Context) ([]entity.Post, error) {
	query := "select id, title, txt from posts"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query select statement: %w", err)
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var (
			id          int64
			title, text string
		)

		err = rows.Scan(&id, &title, &text)
		if err != nil {
			return nil, fmt.Errorf("failed to rows scan: %w", err)
		}

		post := entity.Post{
			ID:    id,
			Title: title,
			Text:  text,
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to rows err: %w", err)
	}
	return posts, nil
}
