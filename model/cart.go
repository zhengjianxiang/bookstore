package model

//购物车
type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
}

//GetTotalCount
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中购物项切片
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中购物项切片
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
