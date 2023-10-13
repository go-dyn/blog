package controller

import (
	"net/http"

	"github.com/go-dyn/blog/app/service"
)

type UserHandler struct {
	BaseHandler
	srv service.UserService
}

var _ http.Handler = (*UserHandler)(nil)

func NewUserHandler() http.Handler {
	return &UserHandler{
		srv: *service.NewUserService(),
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:

		users, err := h.srv.List(0, 10)
		if err != nil {
			h.Response(w, http.StatusInternalServerError, map[string]any{"message": err})
			return
		}
		h.Response(w, http.StatusInternalServerError, map[string]any{"message": "Success", "items": users})
		return
	case http.MethodPost:

		id, err := h.srv.Create(req)
		if err != nil {
			h.Response(w, http.StatusInternalServerError, map[string]any{"message": err})
			return
		}

		item, err := h.srv.Get(id)

		if err != nil {
			h.Response(w, http.StatusInternalServerError, map[string]any{"message": err})
			return
		}

		h.Response(w, http.StatusInternalServerError, map[string]any{"message": "Success", "item": item})
		return
	}
	h.Response(w, http.StatusNotFound, map[string]string{"message": "Not Found"})
}
