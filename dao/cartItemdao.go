package dao

import (
	"goweb/day03/bookstore0612/model"
	"goweb/day03/bookstore0612/utils"
)

//向购物项表插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql语句
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil

}

//根据图书id获取购物项
func GetCartItemByBookID(bookID string) (*model.CartItem, error) {
	//写sql
	sqlStr := "select id,COUNT,amount,cart_id from cart_items where book_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	//创建
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil

}

func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,COUNT,amount,cart_id from cart_items where cart_id = ?"

	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
