package dao

import (
	"fmt"
	"goweb/day03/bookstore0612/model"
	"testing"
)

func TestMain(m *testing.M) {
	// fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	// fmt.Println("测试userdao中的函数")
	// t.Run("验证用户名或密码：", testLogin)
	// t.Run("验证用户名：", testRegist)
	// t.Run("保存用户：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@atguigu.com")
}

func TestBook(t *testing.T) {
	// fmt.Println("测试books")
	// t.Run("测试所有图书", testGetBooks)
	// t.Run("测试添加所有图书", testAddBooks)
	// t.Run("测试删除图书", testDelteBook)
	// t.Run("测试获取一本图书", testGetBook)
	// t.Run("测试更新图书", testUpdateBook)
	// t.Run("测试获取🍁图书", testGetPageBooks)
	// t.Run("测试带分页价格范围图书", testGetPageBooksByPrice)

}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	// fmt.Println(books)
	//遍历所有图书
	for k, v := range books {
		fmt.Printf("第%v本图书信息：%v\n", k+1, v)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	//调用添加图书函数
	AddBook(book)
}

func testDelteBook(t *testing.T) {

	//调用删除图书函数
	DeleteBook("34")
}

func testGetBook(t *testing.T) {

	//调用获取图书函数
	book, _ := GetBookByID("30")
	fmt.Println("获取图书信息：", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      32,
		Title:   "水浒传高配",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	//调用获取图书函数
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("1")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("zong页是：", page.TotalPageNo)
	fmt.Println("总计绿树是：", page.TotalRecord)
	fmt.Println("tushu：")
	for _, v := range page.Books {
		fmt.Println("图书信息", v)
	}
}

func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("3", "10", "30")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("zong页是：", page.TotalPageNo)
	fmt.Println("总计绿树是：", page.TotalRecord)
	fmt.Println("tushu：")
	for _, v := range page.Books {
		fmt.Println("图书信息", v)
	}
}

func TestSession(t *testing.T) {
	fmt.Println("session相关")
	// t.Run("测试添加", testAddSession)
	// t.Run("测试删除", testDeleteSession)
	// t.Run("测试获取Session",testGetSession)
}

func TestCart(t *testing.T) {
	fmt.Println("测试购物车中相关函数")
	// t.Run("测试添加购物车", testAddCart)
	// t.Run("测试根据图书id获取账户购物项", testGetCartItemByBookID)
	// t.Run("测试根据购物车id获取所有购物项", testGetCartItemByCartID)
	t.Run("测试根据用户id获取对应购物车", testGetCartByUserID)
}

func testAddCart(t *testing.T) {
	//设置要买的第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.2,
	}
	//设置第二部
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	var cartItems []*model.CartItem
	//创建购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "001",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "001",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物项
	cart := &model.Cart{
		CartID:    "001",
		CartItems: cartItems,

		UserID: 3,
	}
	AddCart(cart)
}
func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838381438",
		UserName:  "maronf",
		UserID:    5,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}
func testGetSession(t *testing.T) {
	sess, _ := GetSession("05541909-0f2e-4fdd-605c-df530ce70469")
	fmt.Println("Session的信息是：", sess)

}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookID("1")
	fmt.Println("图书id=1", cartItem)
}

func testGetCartItemByCartID(t *testing.T) {
	cartItems, _ := GetCartItemByCartID("001")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是：%v\n", k+1, v)
	}
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(3)
	fmt.Println("id为3用户", cart)
}

// func TestBook(t *testing.T) {
// 	fmt.Println("测试bookdao中的相关函数")
// 	// t.Run("测试获取所有图书", testGetBooks)
// 	// t.Run("测试添加图书", testAddBook)
// 	// t.Run("测试删除图书", testDeleteBook)
// 	// t.Run("测试获取一本图书", testGetBook)
// 	t.Run("测试更新图书", testUpdateBook)
// }

// func testGetBooks(t *testing.T) {
// 	books, _ := GetBooks()
// 	//遍历得到每一本图书
// 	for k, v := range books {
// 		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
// 	}
// }
// func testAddBook(t *testing.T) {
// 	book := &model.Book{
// 		Title:   "三国演义",
// 		Author:  "罗贯中",
// 		Price:   88.88,
// 		Sales:   100,
// 		Stock:   100,
// 		ImgPath: "/static/img/default.jpg",
// 	}
// 	//调用添加图书的函数
// 	AddBook(book)
// }
// func testDeleteBook(t *testing.T) {
// 	//调用删除图书的函数
// 	DeleteBook("34")
// }
// func testGetBook(t *testing.T) {
// 	//调用获取图书的函数
// 	book, _ := GetBookByID("32")
// 	fmt.Println("获取的图书信息是：", book)
// }
// func testUpdateBook(t *testing.T) {
// 	book := &model.Book{
// 		ID:      32,
// 		Title:   "3个女人与105个男人的故事",
// 		Author:  "罗贯中",
// 		Price:   66.66,
// 		Sales:   10000,
// 		Stock:   1,
// 		ImgPath: "/static/img/default.jpg",
// 	}
// 	//调用更新图书的函数
// 	UpdateBook(book)
// }
