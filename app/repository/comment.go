package repository

import "database/sql"

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		db: getInstance(),
	}
}

func (r *CommentRepository) GetById(id int64) *sql.Row {
	row := r.db.QueryRow("select * from comments where id = ? LIMIT ?", id, 1)
	return row
}

func (r *CommentRepository) Create(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("INSERT INTO comments(`name`) VALUES(?)", args)
	return result, err
}

func (r *CommentRepository) Update(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE TABLE comments SET `name`= ?", args)
	return result, err
}

func (r *CommentRepository) Delete(id int64) (sql.Result, error) {
	result, err := r.db.Exec("DELETE FROM comments WHERE `id`= ?", id)
	return result, err
}
