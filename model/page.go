// package model

// //Page 结构
// type Page struct {
// 	Books       []*Book //每页查询出来的图书存放的切片
// 	PageNo      int64   //当前页
// 	PageSize    int64   //每页显示的条数
// 	TotalPageNo int64   //总页数，通过计算得到
// 	TotalRecord int64   //总记录数，通过查询数据库得到
// 	MinPrice    string
// 	MaxPrice    string
// }

// //IsHasPrev 判断是否有上一页
// func (p *Page) IsHasPrev() bool {
// 	return p.PageNo > 1
// }

// //IsHasNext 判断是否有下一页
// func (p *Page) IsHasNext() bool {
// 	return p.PageNo < p.TotalPageNo
// }

package model

type Page struct {
	Books       []*Book //每页查询的图书切片
	PageNo      int64   //当前也
	PageSize    int64   //每页显示条数
	TotalPageNo int64   //总页数 通过计算得到
	TotalRecord int64   //总记录 通过查询得到
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 判断下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetPrevPageNo 获取上一页
// func (p *Page) GetPrevPageNo() int64 {
// 	if p.IsHasPrev() {
// 		return p.PageNo - 1
// 	} else {
// 		return 1
// 	}
// }

//GetNextPageNo 获取下一页
// func (p *Page) GetNextPageNo() int64 {
// 	if p.IsHasNext(){
// 		return p.PageNo+1
// 	}else{
// 		return p.TotalPageNo
// 	}
// }
//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
