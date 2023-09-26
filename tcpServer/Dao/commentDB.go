package Dao

import (
	"entryTask/common/log"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CommentInfo struct {
	ID        uint64 `xorm:"id"`
	Comment   string `xorm:"comment"`
	Uid       string `xorm:"uid"`
	Ctime     uint64 `xorm:"ctime"`
	IsVisible int    `xorm:"is_visible"`
	MessageId uint64 `xorm:"message_id"`
}

func GetCommentListByUid(uid string, page, pageSize int32) ([]*CommentInfo, error) {
	sqlStr := "select id,comment,uid,ctime,is_visible,message_id from comment_tab"
	if len(uid) != 0 {
		sqlStr += fmt.Sprintf(" where uid='%s'", uid)
	}
	if page > 0 && pageSize > 0 {
		sqlStr += fmt.Sprintf(" order by ctime desc limit %d, %d", (page-1)*pageSize, pageSize)
	}

	list := make([]*CommentInfo, 0)

	row, err := db.Query(sqlStr)
	if err != nil {
		log.Log.Errorf("db GetCommentList err:%s", err)
		return nil, err
	}
	for row.Next() {
		m := &CommentInfo{}
		err = row.Scan(&m.ID, &m.Comment, &m.Uid, &m.Ctime, &m.IsVisible, &m.MessageId)
		if err != nil {
			log.Log.Errorf("err:%s", err)
			return nil, err
		}
		list = append(list, m)
	}

	return list, nil
}

func GetCommentListByMessageIds(msgIds []uint64, page, pageSize int32) ([]*CommentInfo, error) {
	sqlStr := "select id,comment,uid,ctime,is_visible,message_id from comment_tab"
	if len(msgIds) > 0 {
		msgIdStr := make([]string, 0)
		for _, elem := range msgIds {
			msgIdStr = append(msgIdStr, strconv.FormatInt(int64(elem), 10))
		}
		sqlStr += fmt.Sprintf(" where message_id IN(%s)", strings.Join(msgIdStr, ","))
	}
	if page > 0 && pageSize > 0 {
		sqlStr += fmt.Sprintf(" order by ctime desc limit %d, %d", (page-1)*pageSize, pageSize)
	}

	list := make([]*CommentInfo, 0)

	row, err := db.Query(sqlStr)
	if err != nil {
		log.Log.Errorf("db GetCommentList err:%s", err)
		return nil, err
	}
	for row.Next() {
		m := &CommentInfo{}
		err = row.Scan(&m.ID, &m.Comment, &m.Uid, &m.Ctime, &m.IsVisible, &m.MessageId)
		if err != nil {
			log.Log.Errorf("err:%s", err)
			return nil, err
		}
		list = append(list, m)
	}

	return list, nil
}

func AddComment(uid, comment string) error {
	t := time.Now().Unix()
	sql := fmt.Sprintf("insert into comment_tab(`comment`,`uid`,`ctime`, `is_visible`,`message_id`) values('%s','%s',%d, %d,'%s')",
		comment, uid, t, 1)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(id uint64) error {
	sql := fmt.Sprintf("delete from comment_tab where id=%d", id)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
