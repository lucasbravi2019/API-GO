package main

import (
	"net/http"
	"pasteleria/api"
	"pasteleria/core"
)

func main() {
	router := core.GetRouter()

	for _, route := range api.GetRoutes() {
		router.Path(route.Path).HandlerFunc(route.Handler).Methods(route.Method)
	}

	http.ListenAndServe("localhost:8080", router)
}
