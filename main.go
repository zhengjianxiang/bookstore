package main

import (
	"goweb/day03/bookstore0612/controller"
	"net/http"
)

func main() {
	//设置静态资源 如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	// 直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", controller.IndexHander)

	//去登陆
	http.HandleFunc("/login", controller.Login)
	//logout
	http.HandleFunc("/logout", controller.Logout)

	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	// 获取所有图书
	// http.HandleFunc("/getBooks", controller.GetBooks)

	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	// 添加图书/getPageBooksByPrice

	// http.HandleFunc("/addBook", controller.AddBook)

	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)

	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)

//添加图书购物车
http.HandleFunc("/addBook2Cart",controller.AddBook2Cart)

	http.ListenAndServe(":8080", nil)
}
