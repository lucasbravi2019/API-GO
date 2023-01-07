package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string `json:"path"`
	Handler http.HandlerFunc
	Method  string `json:"method"`
}

type Routes []Route

func GetRouter() *mux.Router {
	return mux.NewRouter()
}
