package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"AddItemToCard",
		"POST",
		"/cart/{cartId}",
		AddItemToCard,
	},

	Route{
		"Checkout",
		"PATCH",
		"/cart/{cartId}",
		Checkout,
	},

	Route{
		"GetByCartId",
		"GET",
		"/cart/{cartId}",
		GetByCartId,
	},

	Route{
		"GetShoppingCart",
		"POST",
		"/cart",
		GetShoppingCart,
	},

	Route{
		"UpdateItemQuantity",
		"PATCH",
		"/cart/{cartId}/{itemId}",
		UpdateItemQuantity,
	},

}
