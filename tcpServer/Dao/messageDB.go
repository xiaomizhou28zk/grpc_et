package Dao

import (
	"entryTask/common/log"
	"fmt"
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

func GetMessageList(uid string) ([]*MessageInfo, error) {

	sqlStr := "select id,message,image,owner,ctime,mtime,uid from message_tab"
	if len(uid) != 0 {
		sqlStr += fmt.Sprintf("where uid=%s", uid)
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
