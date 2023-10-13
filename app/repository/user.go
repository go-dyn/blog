package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: getInstance(),
	}
}

func (r *UserRepository) List(offset int64, limit int64) (*sql.Rows, error) {
	rows, err := r.db.Query("select * from users where LIMIT ?, ?", offset, limit)
	return rows, err
}

func (r *UserRepository) Get(id int64) *sql.Row {
	row := r.db.QueryRow("select * from users where id = ? LIMIT ?", id, 1)
	return row
}

func (r *UserRepository) Create(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("INSERT INTO users(`name`) VALUES(?)", args)
	return result, err
}

func (r *UserRepository) Update(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE TABLE users SET `name`= ?", args)
	return result, err
}

func (r *UserRepository) Delete(id int64) (sql.Result, error) {
	result, err := r.db.Exec("DELETE FROM users WHERE `id`= ?", id)
	return result, err
}
