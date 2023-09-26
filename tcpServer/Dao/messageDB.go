package Dao

import (
	"entryTask/common/log"
	"fmt"
	"time"
)

type MessageInfo struct {
	ID      uint64 `xorm:"id"`
	Message string `xorm:"message"`
	Image   string `xorm:"image"`
	Owner   string `xorm:"owner"`
	Ctime   uint64 `xorm:"ctime"`
	Mtime   uint64 `xorm:"mtime"`
	Uid     string `xorm:"uid"`
}

func GetMessageList(uid string, page, pageSize int32) ([]*MessageInfo, error) {

	sqlStr := "select id,message,image,owner,ctime,mtime,uid from message_tab"
	if len(uid) != 0 {
		sqlStr += fmt.Sprintf(" where uid='%s'", uid)
	}
	if page > 0 && pageSize > 0 {
		sqlStr += fmt.Sprintf(" order by ctime desc limit %d, %d", (page-1)*pageSize, pageSize)
	}

	msgList := make([]*MessageInfo, 0)

	row, err := db.Query(sqlStr)
	if err != nil {
		log.Log.Errorf("db get userinfo err:%s", err)
		return nil, err
	}
	for row.Next() {
		m := &MessageInfo{}
		err = row.Scan(&m.ID, &m.Message, &m.Image, &m.Owner, &m.Ctime, &m.Mtime, &m.Uid)
		if err != nil {
			log.Log.Errorf("err:%s", err)
			return nil, err
		}
		msgList = append(msgList, m)
	}

	return msgList, nil
}

func AddMessage(uid, msg, userName string) error {
	t := time.Now().Unix()
	sql := fmt.Sprintf("insert into message_tab(`message`,`owner`,`uid`,`ctime`,`mtime`) values('%s','%s','%s',%d,%d)",
		msg, userName, uid, t, t)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func GetMessageCount(uid string) (int32, error) {
	sql := fmt.Sprintf("select count(*) from message_tab")

	if uid != "" {
		sql += fmt.Sprintf(" where uid='%s'", uid)
	}
	var count int32
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Log.Errorf("db GetMessageCount err:%s", err)
		return 0, err
	}
	return count, nil
}

func DeleteMessage(id uint64) error {
	sql := fmt.Sprintf("delete from message_tab where id=%d", id)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
