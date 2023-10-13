package repository

import "database/sql"

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		db: getInstance(),
	}
}

func (r *PostRepository) GetById(id int64) *sql.Row {
	row := r.db.QueryRow("select * from posts where id = ? LIMIT ?", id, 1)
	return row
}

func (r *PostRepository) Create(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("INSERT INTO posts(`name`) VALUES(?)", args)
	return result, err
}

func (r *PostRepository) Update(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE TABLE posts SET `name`= ?", args)
	return result, err
}

func (r *PostRepository) Delete(id int64) (sql.Result, error) {
	result, err := r.db.Exec("DELETE FROM posts WHERE `id`= ?", id)
	return result, err
}
