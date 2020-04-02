package user

import (
	"common-lib/mysql"
	"jaden/we-life/common"
	"jaden/we-life/dao"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
)

type Dao struct {
	dao.BaseDao
}

func NewDao() *Dao {
	return &Dao{}
}

//通过id或者名字获取user信息
func (dao *Dao) GetUserByIdOrName(id int, name string) (*User, *entity.ErrCode) {
	var (
		sql    = `select id,name,birthday,avatar,phone,password,last_login_time,login_count from user where 1=1 `
		params = make([]interface{}, 0)
	)
	if id != 0 {
		params = append(params, id)
		sql += ` and id =? `
	}
	if name != "" {
		params = append(params, name)
		sql += ` and name=? `
	}
	if params == nil || len(params) == 0 {
		return nil, errcode.PARAM_ERROR
	}
	user := new(User)
	err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		rows.Scan(&user.Id, &user.Name, &user.Birthday, &user.Avatar, &user.Phone, &user.Password, &user.LastLoginTime, &user.LoginCount)
	})
	if err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA
		}
		return nil, errcode.SYSTEM_ERROR
	}
	return user, nil
}

//通过user主键更新信息
func (dao Dao) UpdateUserInfoById(user *User) *entity.ErrCode {
	var (
		params = make([]interface{}, 0)
		sql    = `update user set `
	)
	if user == nil {
		return errcode.PARAM_ERROR
	}

	/*  if user.Name != ""{
	        sql += ` name = ?,`
	        params = append(params,user.Name)
	}*/
	if user.LoginCount != 0 {
		sql += ` login_count=?,`
		params = append(params, user.LoginCount)
	}
	if user.LastLoginTime != 0 {
		sql += ` last_login_time=?,`
		params = append(params, user.LastLoginTime)
	}

	if user.Password != "" {
		sql += ` password=?,`
		params = append(params, user.Password)
	}
	if user.Phone != "" {
		sql += `phone=?,`
		params = append(params, user.Phone)
	}
	if user.Avatar != "" {
		sql += `avatar=?,`
		params = append(params, user.Avatar)
	}
	if user.Birthday != 0 {
		sql += `birthday=?,`
		params = append(params, user.Birthday)
	}
	//去掉逗号
	sql = sql[:len(sql)-1]

	sql += ` where id=?`
	params = append(params, user.Id)
	if _, err := dao.DBUpdate(common.DB, sql, params); err != nil {
		return err
	}
	return nil

}

//新增一个用户
func (dao Dao) InsertUser(user *User) *entity.ErrCode {
	var (
		sql    = `insert into user(name,birthday,avator,phone,password) values(?,?,?,?,?)`
		params = make([]interface{}, 0)
	)
	params = append(params, user.Name, user.Birthday, user.Avatar, user.Phone, user.Password)
	if _, err := dao.DBInsert(common.DB, sql, params); err != nil {
		return err
	}
	return nil
}

//查询绑定的另一半的信息
func (dao Dao) QueryBindUser(userId int) (*UserBind, *entity.ErrCode) {
	var (
		sql    = `select user_id ,bind_user_id, is_accepted,unix_timestamp(bind_time) from user_bind where 1=1 `
		params = make([]interface{}, 0)
		result = new(UserBind)
	)
	if userId == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("用户id为空")
	}
	sql += ` and user_id=? and is_accepted=1`
	params = append(params, userId)
	err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		_ = rows.Scan(&result.UserId, &result.BindUserId, &result.IsAccepted, &result.BindTime)
	})
	if err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA.AddMsg("不存在绑定的另一半")
		}
		return nil, err
	}
	return result, nil
}

func (dao Dao) CreateUserBind(bind *UserBind) *entity.ErrCode {
	var (
		sql    = `insert into user_bind(user_id,bind_user_id,is_accepted,bind_time) values(?,?,?,now())`
		params = make([]interface{}, 0)
	)
	if bind == nil {
		return errcode.PARAM_ERROR.AddMsg("绑定用户的信息为空")
	}
	params = append(params, bind.UserId, bind.BindUserId, bind.IsAccepted)
	if _, err := dao.DBInsert(common.DB, sql, params); err != nil {
		return err
	}
	return nil
}

func (dao *Dao) GetRequestBindUserList(userBindId int) ([]*UserBind, *entity.ErrCode) {

	var (
		sql    = `select user_id ,bind_user_id, is_accepted,unix_timestamp(bind_time) from user_bind where 1=1 `
		params = make([]interface{}, 0)
		res    = make([]*UserBind, 0)
	)
	if userBindId == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("用户id为空")
	}
	sql += ` and bind_user_id=?`
	err := dao.DBQuery(common.DB, sql, params, func(rows *mysql.Rows) {
		result := new(UserBind)
		_ = rows.Scan(&result.UserId, &result.BindUserId, &result.IsAccepted, &result.BindTime)
		res = append(res, result)
	})
	if err != nil {
		if err.Code == errcode.DATA_NOT_EXIST.Code {
			return nil, errcode.EMPTY_DATA.AddMsg("不存在绑定的另一半")
		}
		return nil, err
	}

	return res, nil
}

func (dao *Dao) UpdateBindStatus(userId int) *entity.ErrCode {
	var (
		params = make([]interface{}, 0)
		sql    = `update user_bind set `
	)
	if userId == 0 {
		return errcode.PARAM_ERROR
	}

	if userId != 0 {
		sql += ` is_accepted=? where user_id =?`
		params = append(params, 1)
	}

	if _, err := dao.DBUpdate(common.DB, sql, params); err != nil {
		return err
	}
	return nil
}
