package rest_api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reqyt.run/core_api/config"
	"strconv"
	"strings"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RemoteAddr, " ", r.Method, " ", r.RequestURI, r.Context().Value("user"))
		next.ServeHTTP(w, r)
	})
}

func userIDFromAuthorizationHeader(h string) (id int, ok bool) {
	if h == "" { return 0, false }
	split := strings.Split(h, " ")
	if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
		return 0, false
	}
	if userID, err := strconv.Atoi(split[1]); err != nil {
		return 0, false
	} else {
		return userID, true
	}
}

type key string
const UserIDKey = key("userID")

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		userID, ok := userIDFromAuthorizationHeader(authHeader)
		if ok {
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Handler404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	resp, _ := json.Marshal(map[string]interface{}{
		"error": "No matching route.",
	})
	w.Write(resp)
}

func Handler403(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	resp, _ := json.Marshal(map[string]interface{}{
		"error": "Not allowed.",
	})
	w.Write(resp)
}

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	AddUserHandlers(r.PathPrefix("/users").Subrouter())
	if config.Config.LogRequests {
		r.Use(jsonMiddleware)
		r.Use(loggingMiddleware)
		r.Use(authMiddleware)
	}
	r.NotFoundHandler = http.HandlerFunc(Handler404)
	r.MethodNotAllowedHandler = http.HandlerFunc(Handler404)
	return r
}