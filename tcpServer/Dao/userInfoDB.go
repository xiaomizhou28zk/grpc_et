package Dao

import (
	"database/sql"
	"entryTask/common/log"
	"entryTask/httpServer/common"
	"entryTask/tcpServer/config"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type UserInfo struct {
	Uid      string `xorm:"uid"`
	Nick     string `xorm:"nick"`
	Picture  string `xorm:"picture"`
	Password string `xorm:"password"`
}

// InitDB 初始化数据库
func InitDB() (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.DBUserName, config.Config.DBPassword,
		config.Config.DBIP, config.Config.DBPort, config.Config.DBName)
	db, err = sql.Open(config.Config.DBDriverName, dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Log.Errorf("db err:%s", err)
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}

// QueryUserInfo 查询用户信息
func QueryUserInfo(uid string) (UserInfo, error) {

	sqlStr := "select uid, nick, picture, password from " + getTableName(uid) + " where uid=?"

	var u UserInfo
	err := db.QueryRow(sqlStr, uid).Scan(&u.Uid, &u.Nick, &u.Picture, &u.Password)
	if err != nil {
		log.Log.Errorf("db get userinfo err:%s", err)
		return u, err
	}

	return u, nil
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(uid, nick, picture string) error {

	sqlStr := "update " + getTableName(uid) + " set nick=?,picture=? where uid=?"

	ret, err := db.Exec(sqlStr, nick, picture, uid)
	if err != nil {
		log.Log.Errorf("update userinfo err:%s", err)
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Log.Errorf("get RowsAffected err:%s", err)
		return err
	}
	if n != 1 {
		return fmt.Errorf("RowsAffected not one")
	}
	return nil
}

// getTableName 获取表名
func getTableName(uid string) string {
	return "user_info_" + strconv.Itoa(int(common.BKDRHash(uid)%100))
}
