package news

import (
	"common-lib/mysql"
	"jaden/we-life/common"
	"jaden/we-life/dao"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
	"jaden/we-life/util"
	"time"
)

func (dao Dao) GetNewsByUserIDs(userId []int) ([]*News, *entity.ErrCode) {
	var (
		sql    = `select id ,content,unix_timestamp(create_time),address,user_id,is_delete from news where 1=1 `
		params = make([]interface{}, 0)
		result = make([]*News, 0)
	)
	if len(userId) == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("没有传入userIds")
	}
	sql += util.AndIn("user_id", userId)
	params = append(params, userId)
	sql += ` and is_delete=0 order by create_time desc`
	params = append(params, userId)
	if err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		n := new(News)
		_ = rows.Scan(&n.Id, &n.Content, &n.CreateTime, &n.Address, &n.UserId, &n.IsDelete)
		result = append(result, n)
	}); err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA
		}
		return nil, err
	}
	return result, nil
}

//新增一条动态
func (dao Dao) InsertNew(userId int, content string, address string) (int, *entity.ErrCode) {
	var (
		sql = `insert into news (content,create_time,address,user_id) values(?,?,?,?)`
	)
	if userId == 0 {
		return 0, errcode.PARAM_ERROR
	}
	lastInsertId, err := dao.DBInsert(common.DB, sql, []interface{}{content, util.TimestampToDateTime(time.Now().Unix()), address, userId})
	if err != nil {
		return 0, err
	}
	return int(lastInsertId), nil
}

func (dao Dao) DeleteNew(newId int) *entity.ErrCode {
	var (
		sql = `update news set is_delete =1 where id = ?`
	)
	if newId == 0 {
		return errcode.PARAM_ERROR
	}
	_, err := dao.DBUpdate(common.DB, sql, []interface{}{newId})
	if err != nil {
		return err
	}
	return nil
}

func (dao Dao) UpdateDirTitle(dirId int, title string) *entity.ErrCode {
	var (
		sql = `update directory_info set dir_name=? where id = ?`
	)
	if dirId == 0 {
		return errcode.PARAM_ERROR
	}
	_, err := dao.DBUpdate(common.DB, sql, []interface{}{title, dirId})
	if err != nil {
		return err
	}
	return nil
}

type Dao struct {
	dao.BaseDao
}

func NewDao() *Dao {
	return &Dao{}
}
