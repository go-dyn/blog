package repository

import "database/sql"

type TagRepository struct {
	db *sql.DB
}

func NewTag() *TagRepository {
	return &TagRepository{
		db: getInstance(),
	}
}

func (r *TagRepository) GetById(id int64) *sql.Row {
	row := r.db.QueryRow("select * from tags where id = ? LIMIT ?", id, 1)
	return row
}

func (r *TagRepository) Create(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("INSERT INTO tags(`name`) VALUES(?)", args)
	return result, err
}

func (r *TagRepository) Update(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE TABLE tags SET `name`= ?", args)
	return result, err
}

func (r *TagRepository) Delete(id int64) (sql.Result, error) {
	result, err := r.db.Exec("DELETE FROM tags WHERE `id`= ?", id)
	return result, err
}
