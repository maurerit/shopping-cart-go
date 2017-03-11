package model

type ShoppingCart struct {
	ShoppingCartId int64 `gorm:"type:bigint(20);primary_key;AUTO_INCREMENT;column:shopping_cart_id" json:"shoppingCartId"`
	CustomerId int64 `gorm:"type:bigint(20);column:customer_id" json:"customerId"`
	Status int `gorm:"type:integer;column:status" json:"status"`
	Items[] ShoppingCartItem `gorm:"-" json:"items"`
}

type ShoppingCartItem struct {
	ShoppingCartId int64 `gorm:"type:bigint(20);primary_key;column:shopping_cart_id" json:"shoppingCartId"`
	ItemId int64 `gorm:"type:bigint(20);primary_key;column:item_id" json:"itemId"`
	Quantity int64 `gorm:"type:bigint(20);column:quantity" json:"quantity"`
	Price float64 `gorm:"type:double;column:price" json:"price"`
	Status int `gorm:"type:integer;column:status" json:"status"`
}

func (s ShoppingCart) TableName() string {
	return "shopping_cart"
}

func (si ShoppingCartItem) TableName() string {
	return "shopping_cart_item"
}
