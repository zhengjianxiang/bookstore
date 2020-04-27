package dao

import (
	"goweb/day03/bookstore0612/model"
	"goweb/day03/bookstore0612/utils"
)

//AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取购物车中购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, cartItem := range cartItems {
		//保存购物项
		AddCartItem(cartItem)
	}
	return nil
}

//根据用户id查询购物车
func GetCartByUserID(userID int) (*model.Cart,error) {
//sql
sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
//执行
row := utils.Db.QueryRow(sqlStr,userID)

cart := &model.Cart{}
err := row.Scan(&cart.CartID,&cart.TotalCount,&cart.TotalAmount,&cart.UserID)
if err != nil {
	return nil,err
}
//获取当前购物车所有购物项
cartItems, _ := GetCartItemByCartID(cart.CartID)
//将所有购物车设置到购物车中
cart.CartItems = cartItems
return cart,nil

}