package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-dyn/blog/app/model"
	"github.com/go-dyn/blog/app/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repository: &repository.UserRepository{},
	}
}

func (u *UserService) List(offset int64, limit int64, args ...any) ([]model.User, error) {
	users := []model.User{}
	rows, err := u.repository.List(offset, limit)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var a model.User
		err = rows.Scan(&a.ID, &a.Email, &a.Name, &a.CreatedAt)
		if err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, a)
	}

	return users, err
}

func (u *UserService) Get(id int64) (model.User, error) {
	user := model.User{}
	row := u.repository.Get(id)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)

	return user, err
}

func (u *UserService) Create(req *http.Request) (int64, error) {
	var createReq model.CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&createReq); err != nil {
		return 0, err
	}
	hash, err := NewSecurityService().Hash(createReq.Password)

	if err := json.NewDecoder(req.Body).Decode(&createReq); err != nil {
		return 0, err
	}

	now := time.Now().Unix()
	item := model.User{
		Name:      createReq.Name,
		Email:     createReq.Email,
		Password:  hash,
		CreatedAt: uint64(now),
		UpdatedAt: uint64(now),
	}
	result, err := u.repository.Create(item)

	if err := json.NewDecoder(req.Body).Decode(&createReq); err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	return id, err
}
