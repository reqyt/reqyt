package rest_api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reqyt.run/core_api/models"
	"strconv"
	"time"
)

type LiteUser struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
func UserListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var users []*LiteUser
	_ = models.Users().BindG(ctx, &users)
	usersString, _ := json.Marshal(users)
	w.Write(usersString)
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	m := &models.User{}
	CreateHandler(m, w, r)
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if id != r.Context().Value(UserIDKey) {
		Handler403(w, r)
		return
	}
	u, err := models.FindUserG(context.Background(), int64(id))
	HandleFind(u, err, w)
}

func AddUserHandlers(s *mux.Router) {
	s.HandleFunc("/", UserListHandler).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", UserGetHandler).Methods(http.MethodGet)
	s.HandleFunc("/", UserCreateHandler).Methods(http.MethodPost)
}
