package postgres

import (
	"GoNews/pkg/storage"
	"database/sql"

	_ "github.com/lib/pq"
)

// Store представляет собой тип хранилища для PostgreSQL.
type Store struct {
	db *sql.DB
}

// New создаёт новое хранилище.
func New(connectionString string) (*Store, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query("SELECT id, title, content, author_id, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []storage.Post
	for rows.Next() {
		var post storage.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Store) AddPost(post storage.Post) error {
	_, err := s.db.Exec("INSERT INTO posts (title, content, author_id, created_at) VALUES ($1, $2, $3, $4)",
		post.Title, post.Content, post.AuthorID, post.CreatedAt)
	return err
}

func (s *Store) UpdatePost(post storage.Post) error {
	_, err := s.db.Exec("UPDATE posts SET title = $1, content = $2, author_id = $3, created_at = $4 WHERE id = $5",
		post.Title, post.Content, post.AuthorID, post.CreatedAt, post.ID)
	return err
}

func (s *Store) DeletePost(post storage.Post) error {
	_, err := s.db.Exec("DELETE FROM posts WHERE id = $1", post.ID)
	return err
}
