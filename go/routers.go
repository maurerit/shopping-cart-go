package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"AddItemToCard",
		"POST",
		"//cart/{cartId}",
		AddItemToCard,
	},

	Route{
		"Checkout",
		"PATCH",
		"//cart/{cartId}",
		Checkout,
	},

	Route{
		"GetByCartId",
		"GET",
		"//cart/{cartId}",
		GetByCartId,
	},

	Route{
		"GetShoppingCart",
		"POST",
		"//cart",
		GetShoppingCart,
	},

	Route{
		"UpdateItemQuantity",
		"PATCH",
		"//cart/{cartId}/{itemId}",
		UpdateItemQuantity,
	},

}