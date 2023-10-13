package repository

import "database/sql"

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: getInstance(),
	}
}

func (r *CategoryRepository) GetById(id int64) *sql.Row {
	row := r.db.QueryRow("select * from categories where id = ? LIMIT ?", id, 1)
	return row
}

func (r *CategoryRepository) Create(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("INSERT INTO categories(`name`) VALUES(?)", args)
	return result, err
}

func (r *CategoryRepository) Update(args ...any) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE TABLE categories SET `name`= ?", args)
	return result, err
}

func (r *CategoryRepository) Delete(id int64) (sql.Result, error) {
	result, err := r.db.Exec("DELETE FROM categories WHERE `id`= ?", id)
	return result, err
}
