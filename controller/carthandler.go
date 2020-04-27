package controller

import (
	_ "fmt"
	"goweb/day03/bookstore0612/dao"
	"goweb/day03/bookstore0612/model"
	"goweb/day03/bookstore0612/utils"
	"net/http"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	// fmt.Println("要添加图书id",bookId)
	//根据图书id 获取图书信息
	book, _ := dao.GetBookByID(bookID)
	//判断是否登录
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//判断购物车中是否有当前用户购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		//已经有购物车

	} else {
		//创建购物项
		//1创建购物项
		//生成购物车id
		cartID := utils.CreateUUID()
		cart := &model.Cart{
			CartID: cartID,
			UserID: userID,
		}
		//2.创建购物车购物项
		var cartItems []*model.CartItem

		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cartID,
		}
		//将购物项添加到切片中
		cartItems = append(cartItems, cartItem)
		//将切片设置到cart中
		cart.CartItems = cartItems
		// 将购物车保存到数据库中
		dao.AddCart(cart)
	}
}
