package service

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/maurerit/shopping-cart-go/cartservice/model"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os"
)

var DB gorm.DB

func GetShoppingCart(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	customerId, _ := strconv.ParseInt(string(body), 10, 64)

	var cart model.ShoppingCart
	DB.Where("shopping_cart_id = ? and status = ?", customerId, 0).FirstOrCreate(&cart)

	var items[] model.ShoppingCartItem
	DB.Where("shopping_cart_id = ?", cart.ShoppingCartId).Find(&items)
	cart.Items = items

	data, _ := json.Marshal(cart)

	CommonHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetByCartId(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountId' path parameter from the mux map
	var shoppingCartId = mux.Vars(r)["cartId"]

	var cart model.ShoppingCart
	DB.First(&cart, "shopping_cart_id = ?", shoppingCartId)

	CommonHeader(w)
	if &cart != nil {
		var items []model.ShoppingCartItem
		DB.Where("shopping_cart_id = ?", cart.ShoppingCartId).Find(&items)

		cart.Items = items

		data, _ := json.Marshal(cart)

		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func AddItemToCard(w http.ResponseWriter, r *http.Request) {
	CommonHeader(w)

	//Because no robust framework exists I have to do all of this myself :(
	//First deserialize the item that should have been passed in
	var item model.ShoppingCartItem
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&item)
	defer r.Body.Close()

	//BGN Temp Code
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(item)
	//END Temp Code

	//Then check that it is something and no error happened decoding it
	//At least I hope that's what this is doing... I should probably read the language spec... but meh
	if &item != nil && err != nil {
		var shoppingCartId = mux.Vars(r)["cartId"]

		var cart model.ShoppingCart
		DB.First(&cart, "shopping_cart_id = ?", shoppingCartId)

		//Now search the db for an item with the given item id and cart id...
		//if we find one, we return a 409
		var items []model.ShoppingCartItem
		//TODO: Optimize this so that it's a count and not returning anything...
		//I have to query for ALL the items later anyways...
		DB.Where("shopping_cart_id = ? and item_id = ?", shoppingCartId, item.ItemId).Find(&items)

		//BGN Temp Code
		enc.Encode(items)
		//END Temp Code

		if len(items) > 0 {
			w.WriteHeader(http.StatusConflict)
		} else {
			DB.Where("shopping_cart_id = ?", shoppingCartId).Find(&items)

			item.ShoppingCartId = cart.ShoppingCartId

			cart.Items = items
			cart.Items = append(cart.Items, item)

			DB.Save(&item)

			//The below is definitely going to be it's own function
			//For now... right it again and again so you get sick and tired of it, good practice ;)
			data, _ := json.Marshal(cart)

			w.Header().Set("Content-Length", strconv.Itoa(len(data)))
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	}
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	var shoppingCartId = mux.Vars(r)["cartId"]

	var cart model.ShoppingCart
	DB.First(&cart, "shopping_cart_id = ?", shoppingCartId)

	CommonHeader(w)
	if &cart != nil {
		//Be inefficient here... because my python and java versions were as well (I was lazy so being lazy here)
		cart.Status = 1
		DB.Save(&cart)

		var items []model.ShoppingCartItem
		DB.Where("shopping_cart_id = ?", cart.ShoppingCartId).Find(&items)

		for _,element := range items {
			element.Status = 1
			DB.Save(&element)
		}


		w.WriteHeader(http.StatusOK)
	} else {

	}
}

func UpdateItemQuantity(w http.ResponseWriter, r *http.Request) {
	CommonHeader(w)
	w.WriteHeader(http.StatusOK)
}

func CommonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

