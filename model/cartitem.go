package model

//CartItem
type CartItem struct {
	CartItemID int64   //购物项ID
	Book       *Book   //购物项中图书信息
	Count      int64   //购物项中图书数量
	Amount     float64 //购物项图书金额小计
	CartID     string  //当前购物项属于哪个购物车

}

func (cartItem *CartItem) GetAmount() float64 {
	//获取当前购物项中图书价格
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price

}
