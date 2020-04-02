package photo

import (
	"common-lib/mysql"
	"fmt"
	"jaden/we-life/common"
	"jaden/we-life/dao"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
	"jaden/we-life/util"
	"strings"
	"time"
)

type Dao struct {
	dao.BaseDao
}

func NewDao() *Dao {
	return &Dao{}
}

func (dao Dao) DeleteFilesByNewId(newId int) *entity.ErrCode {
	var (
		sql = `update file_info set is_delete = 1 where new_id=?`
	)
	_, err := dao.DBUpdate(common.DB, sql, []interface{}{newId})
	if err != nil {
		return err
	}
	return nil
}

//获取目录列表
func (dao Dao) GetDirList(userId int) ([]*Dir, *entity.ErrCode) {
	if userId == 0 {
		return nil, errcode.PARAM_ERROR
	}
	var (
		sql    = `select id,dir_name,dir_file_num,UNIX_TIMESTAMP(create_time),user_id,dir_type,is_delete from directory_info where 1=1 `
		params = make([]interface{}, 0)
		result = make([]*Dir, 0)
	)
	sql += `and user_id=? and is_delete = 0`
	params = append(params, userId)
	if err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		res := new(Dir)
		_ = rows.Scan(&res.Id, &res.DirectionName, &res.DirFileCount, &res.CreateTime, &res.UserId, &res.DirType, &res.IsDelete)
		result = append(result, res)
	}); err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA
		}
		return nil, err
	}
	return result, nil
}

//通过newId获取对应的文件
func (dao Dao) GetFilesByNewId(newId int) ([]*FileInfo, *entity.ErrCode) {
	var (
		sql    = `select id,name,path,view_count,unix_timestamp(upload_time),is_delete,directory_id,user_id,new_id from file_info where 1=1`
		params = make([]interface{}, 0)
		result = make([]*FileInfo, 0)
	)

	if newId == 0 {
		return nil, errcode.PARAM_ERROR
	}
	sql += ` and new_id=? and is_delete=0`
	params = append(params, newId)
	err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		fi := new(FileInfo)
		rows.Scan(&fi.Id, &fi.Name, &fi.Path, &fi.ViewCount, &fi.UploadTime, &fi.IsDelete, &fi.DirectoryId, &fi.UserId, &fi.NewId)
		result = append(result, fi)
	})
	if err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA
		}
		return nil, err
	}
	return result, nil
}

//获取用户所有文件总数
func (dao Dao) GetFileCountByDirType(newIds ...int) (int, *entity.ErrCode) {
	var (
		sql    = `select count(id) from file_info where 1=1 and is_delete = 0 `
		params = make([]interface{}, 0)
		result = 0
	)

	sql += util.AndIn("new_id", newIds)
	for _, id := range newIds {
		params = append(params, id)
	}
	if err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		_ = rows.Scan(&result)
	}); err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return 0, errcode.EMPTY_DATA
		}
		return 0, err
	}
	return result, nil

}

//获取文件目录列表
func (dao Dao) GetFileInfoList(directoryId int) ([]*FileInfo, *entity.ErrCode) {
	var (
		sql    = `select id,name,path,view_count,upload_time,is_delete,directory_id,user_id,new_id from file_info where 1=1`
		params = make([]interface{}, 0)
		result = make([]*FileInfo, 0)
	)

	if directoryId == 0 {
		return nil, errcode.PARAM_ERROR
	}
	sql += ` and directory_id=? and is_delete=0`
	params = append(params, directoryId)
	err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		fi := new(FileInfo)
		rows.Scan(&fi.Id, &fi.Name, &fi.Path, &fi.ViewCount, &fi.UploadTime, &fi.IsDelete, &fi.DirectoryId, &fi.UserId, &fi.NewId)
		result = append(result, fi)
	})
	if err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA
		}
		return nil, err
	}
	return result, nil
}

//新增目录
func (dao Dao) InsertDir(dirName string, userId int, dirType int) *entity.ErrCode {
	var (
		sql = `insert into directory_info (dir_name,user_id,dir_type) values(?,?,?)`
	)
	if userId == 0 || dirName == "" {
		return errcode.PARAM_ERROR
	}
	_, err := dao.DBInsert(common.DB, sql, []interface{}{dirName, userId, dirType})
	if err != nil {
		return err
	}
	return nil
}

//新增文件
func (dao Dao) InsertFiles(data *PublishEntity, newId int) *entity.ErrCode {

	sqlStr := `insert into file_info (name,path,upload_time,directory_id,user_id,new_id) values(?,?,?,?,?,?)`

	params := make([]interface{}, 0)

	if data == nil {
		return errcode.PARAM_ERROR
	}
	str := strings.Repeat(",(?,?,?,?,?,?)", len(data.UploadFileEntity)-1)

	sqlStr += str
	fmt.Println(sqlStr)
	for _, info := range data.UploadFileEntity {
		params = append(params, info.FileName, info.FilePath, util.TimestampToDateTime(time.Now().Unix()), data.DirId, data.UserId, newId)
	}
	_, err := dao.DBInsert(common.DB, sqlStr, params)
	if err != nil {
		return err
	}
	return nil
}
