package dao

import (
	"goweb/day03/bookstore0612/model"
	"goweb/day03/bookstore0612/utils"
	"net/http"
)

//AddSession 向数据库中添加session
func AddSession(sess *model.Session) error {
	//sql 语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	//写sql语句
	sqlStr := "delete from sessions where session_id=?"
	//执行sql语句
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionById 根据Session的ID从数据库中查询session
func GetSession(sessID string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建session
	sess := &model.Session{}
	//扫描数据库中的字段为Session
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil

}

//判断用户是否登陆 false没用登录 true 已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据cookie的name获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库查找与之对应session
		session, _ := GetSession(cookieValue)

		if session.UserID > 0 {
			//已经登录
			return true, session

		}
	}
	//没用问题
	return false, nil

}
