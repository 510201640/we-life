package dao

import (
	"common-lib/mysql"
	"database/sql"
	"fmt"
	"jaden/we-life/common"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
)

type DBSession interface {
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type BaseDao struct {
}

func NewBaseDao() *BaseDao {
	baseDao := new(BaseDao)
	return baseDao
}

func (b BaseDao) DBQuery(dbSession DBSession, sqlStr string, arg []interface{}, handler func(rows *mysql.Rows)) *entity.ErrCode {
	var err error

	stmt, err := dbSession.Prepare(sqlStr)
	if err != nil {
		b.logDBError(sqlStr, arg, err)
		return errcode.DATA_ERROR.FromError(err)
	}
	defer stmt.Close()

	rs, err := stmt.Query(arg...)
	if err != nil {
		b.logDBError(sqlStr, arg, err)
		return errcode.DATA_ERROR.FromError(err)
	}

	defer rs.Close()
	var existFlag = false
	for rs.Next() {
		existFlag = true
		handler(rs)
	}
	if !existFlag {
		return errcode.DATA_NOT_EXIST
	}
	return nil
}

func (b BaseDao) DBUpdate(dbSession DBSession, sqlStr string, arg []interface{}) (effectCount int64, errCode *entity.ErrCode) {

	var err error
	var rs sql.Result
	var stmt *sql.Stmt

	defer func() {
		b.logDBError(sqlStr, arg, err)
	}()

	stmt, err = dbSession.Prepare(sqlStr)

	if err != nil {

		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	defer stmt.Close()
	rs, err = stmt.Exec(arg...)
	if err != nil {
		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	effectCount, err = rs.RowsAffected()
	if err != nil {
		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	return
}

func (b BaseDao) DBInsert(dbSession DBSession, sqlStr string, arg []interface{}) (lastId int64, errCode *entity.ErrCode) {

	var err error
	defer func() {
		b.logDBError(sqlStr, arg, err)
	}()

	stmt, err := dbSession.Prepare(sqlStr)
	if err != nil {
		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	defer stmt.Close()
	rs, err := stmt.Exec(arg...)
	if err != nil {
		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	lastId, err = rs.LastInsertId()
	if err != nil {
		errCode = errcode.DATA_ERROR.FromError(err)
		return
	}
	return
}

func (b BaseDao) logDBError(sqlStr string, arg []interface{}, err error) {
	if err == nil {
		return
	}
	common.Logger.Error(fmt.Sprint(arg), err.Error(), sqlStr)
}

func (BaseDao) SqlParam(param ...interface{}) []interface{} {
	return param
}
