package Dao

import (
	"entryTask/common/log"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ReplyInfo struct {
	ID        uint64 `xorm:"id"`
	CommentId uint64 `xorm:"comment_id"`
	Reply     string `xorm:"reply"`
	Uid       string `xorm:"uid"`
	Ctime     uint64 `xorm:"ctime"`
	IsVisible int    `xorm:"is_visible"`
	ToUid     string `xorm:"to_uid"`
}

func GetReplyListByUid(uid string, page, pageSize int32) ([]*ReplyInfo, error) {
	sqlStr := "select id,comment_id,reply,uid,ctime,is_visible,to_uid from reply_tab"
	if len(uid) != 0 {
		sqlStr += fmt.Sprintf(" where uid='%s'", uid)
	}
	if page > 0 && pageSize > 0 {
		sqlStr += fmt.Sprintf(" order by ctime desc limit %d, %d", (page-1)*pageSize, pageSize)
	}

	list := make([]*ReplyInfo, 0)

	row, err := db.Query(sqlStr)
	if err != nil {
		log.Log.Errorf("db GetReplyListByUid err:%s", err)
		return nil, err
	}
	for row.Next() {
		m := &ReplyInfo{}
		err = row.Scan(&m.ID, &m.CommentId, &m.Reply, &m.Uid, &m.Ctime, &m.IsVisible)
		if err != nil {
			log.Log.Errorf("err:%s", err)
			return nil, err
		}
		list = append(list, m)
	}

	return list, nil
}

func GetReplyListByCommentIds(commentIds []uint64, page, pageSize int32) ([]*ReplyInfo, error) {
	sqlStr := "select id,comment_id,reply,uid,ctime,is_visible from reply_tab"
	if len(commentIds) > 0 {
		commentIdStr := make([]string, 0)
		for _, elem := range commentIds {
			commentIdStr = append(commentIdStr, strconv.FormatInt(int64(elem), 10))
		}
		sqlStr += fmt.Sprintf(" where comment_id IN(%s)", strings.Join(commentIdStr, ","))
	}
	if page > 0 && pageSize > 0 {
		sqlStr += fmt.Sprintf(" order by ctime desc limit %d, %d", (page-1)*pageSize, pageSize)
	}

	list := make([]*ReplyInfo, 0)

	row, err := db.Query(sqlStr)
	if err != nil {
		log.Log.Errorf("db GetReplyListByCommentId err:%s", err)
		return nil, err
	}
	for row.Next() {
		m := &ReplyInfo{}
		err = row.Scan(&m.ID, &m.CommentId, &m.Reply, &m.Uid, &m.Ctime, &m.IsVisible, &m.ToUid)
		if err != nil {
			log.Log.Errorf("err:%s", err)
			return nil, err
		}
		list = append(list, m)
	}

	return list, nil
}

func AddReply(uid, reply, toUid string, commentId uint64) error {
	t := time.Now().Unix()
	sql := fmt.Sprintf("insert into reply_tab(`comment_id`, `reply`,`uid`,`ctime`, `is_visible`, `to_uid`) values(%d, '%s','%s',%d, %d, '%s')",
		commentId, reply, uid, t, 1, toUid)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReply(id uint64) error {
	sql := fmt.Sprintf("delete from reply_tab where id=%d", id)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
