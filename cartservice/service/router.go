package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/maurerit/shopping-cart-go/cartservice/logging"
	"fmt"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logging.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		fmt.Println(route.Name + "," + route.Pattern)
	}

	return router
}
