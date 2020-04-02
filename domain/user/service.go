package user

import (
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
	"jaden/we-life/util"
	"time"
)

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Birthday      int    `json:"birthday"`
	Avatar        string `json:"avatar"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	LastLoginTime int    `json:"lastLoginTime"`
	LoginCount    int    `json:"loginCount"`
}

type UserDesc struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday int    `json:"birthday"`
	Avatar   string `json:"avatar"`
}

type UserDetailInfo struct {
	Name          string       `json:"name"`
	Birthday      int          `json:"birthday"`
	Avator        string       `json:"avator"`
	Phone         string       `json:"phone"`
	LoginCount    int          `json:"loginCount"`
	LastLoginTime int          `json:"lastLoginTime"`
	BindUserInfo  BindUserInfo `json:"bindUserInfo"`
}
type BindUserInfo struct {
	Name          string `json:"name"`
	Birthday      int    `json:"birthday"`
	Avator        string `json:"avator"`
	Phone         string `json:"phone"`
	LoginCount    int    `json:"loginCount"`
	LastLoginTime int    `json:"lastLoginTime"`
}

type UserInfo struct {
	Session string `json:"session"`
}

type IService interface {
	GetUserDetailInfoById(id int) (*UserDetailInfo, *entity.ErrCode)

	GetUserByIdOrName(id int, name string) (*User, *entity.ErrCode)

	UpdateUserInfoById(user *User) *entity.ErrCode

	InsertUser(user *User) *entity.ErrCode

	QueryBindUser(userId int) (*User, *entity.ErrCode)

	Login(id int, password string) (*UserInfo, *entity.ErrCode)

	RequestBindUser(userId, bindUserId int) *entity.ErrCode

	UserBindListRequest(userId int) ([]*UserDesc, *entity.ErrCode)

	AgreeBindUser(userId, bindUserID int) *entity.ErrCode
}

type Service struct {
	Dao *Dao
}

func NewService() IService {
	return &Service{
		Dao: NewDao(),
	}
}

//查询请求绑定用户userBindid的用户列表
//用户产生绑定关系后，不可以查询此功能
func (s Service) UserBindListRequest(userBindId int) ([]*UserDesc, *entity.ErrCode) {
	if userBindId == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("用户id为空")
	}
	userBindList, err := s.Dao.GetRequestBindUserList(userBindId)
	if err != nil {
		return nil, err
	}
	result := make([]*UserDesc, 0)
	for _, userBind := range userBindList {
		if userBind.IsAccepted == 1 {
			return nil, errcode.SYSTEM_ERROR.AddMsg("已经存在绑定用户了，不可以查询此功能")
		}
		user, err := s.Dao.GetUserByIdOrName(userBind.BindUserId, "")
		if err != nil {
			return nil, err
		}
		ud := &UserDesc{
			Id:       user.Id,
			Name:     user.Name,
			Birthday: user.Birthday,
			Avatar:   user.Avatar,
		}
		result = append(result, ud)
	}
	return result, nil

	return nil, nil
}

func (s Service) AgreeBindUser(userId, bindUserId int) *entity.ErrCode {
	if userId == 0 || bindUserId == 0 {
		return errcode.PARAM_ERROR
	}
	//查询该用户是否收到邀请
	requestBindList, err := s.Dao.GetRequestBindUserList(userId)
	if err != nil {
		return err
	}
	//查询对方是否已经绑定
	_, err = s.Dao.QueryBindUser(bindUserId)
	if err != nil {
		if err.Code != errcode.EMPTY_DATA.Code {
			return err
		}
	}

	isRecBindRequest := false
	for _, userBind := range requestBindList {
		if userBind.UserId == bindUserId {
			isRecBindRequest = true
		}
	}
	if !isRecBindRequest {
		return errcode.DATA_ERROR.AddMsg("没有收到该用户邀请")
	}
	//更新对方的绑定状态
	if err := s.Dao.UpdateBindStatus(bindUserId); err != nil {
		return err
	}
	//插入我的绑定装填
	userBind := &UserBind{
		UserId:     userId,
		BindUserId: bindUserId,
		IsAccepted: 1,
		BindTime:   time.Now().Unix(),
	}
	err = s.Dao.CreateUserBind(userBind)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) RequestBindUser(userId, bindUserId int) *entity.ErrCode {
	if userId == 0 || bindUserId == 0 {
		return errcode.PARAM_ERROR
	}
	//查询该用户是否绑定了另一个用户
	_, err := s.Dao.QueryBindUser(userId)
	if err != nil {
		if err.Code != errcode.EMPTY_DATA.Code {
			return entity.NewErrCode(1001, "该用户已经被绑定~")
		}
	}
	_, err = s.Dao.GetUserByIdOrName(bindUserId, "")
	if err != nil {
		return err.AddMsg("请求绑定的用户不存在")
	}
	//添加到userBind表信息中
	userBind := &UserBind{
		UserId:     userId,
		BindUserId: bindUserId,
		IsAccepted: 0,
	}
	if err := s.Dao.CreateUserBind(userBind); err != nil {
		return err
	}
	return nil
}

//登录后返回session，请求的时候需要携带session信息
func (s Service) Login(id int, password string) (*UserInfo, *entity.ErrCode) {
	if id == 0 || password == "" {
		return nil, errcode.PARAM_ERROR
	}
	user, err := s.Dao.GetUserByIdOrName(id, "")
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errcode.SESSION_CHECK_ERROR
	}
	session := util.GenSession()
	util.SaveSession(session)

	return &UserInfo{Session: session}, nil
}

func (s Service) GetUserDetailInfoById(id int) (*UserDetailInfo, *entity.ErrCode) {
	if id == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("id为空")
	}
	majorUser, err := s.GetUserByIdOrName(id, "")
	if err != nil {
		return nil, err
	}
	bindUser, err := s.QueryBindUser(id)
	if err != nil {
		return nil, err
	}
	userDetailInfo := &UserDetailInfo{
		Name:          majorUser.Name,
		Birthday:      majorUser.Birthday,
		Avator:        majorUser.Avatar,
		Phone:         majorUser.Phone,
		LoginCount:    majorUser.LoginCount,
		LastLoginTime: majorUser.LastLoginTime,
		BindUserInfo: BindUserInfo{
			Name:          bindUser.Name,
			Birthday:      bindUser.Birthday,
			Avator:        bindUser.Avatar,
			Phone:         bindUser.Phone,
			LoginCount:    bindUser.LoginCount,
			LastLoginTime: bindUser.LastLoginTime,
		},
	}
	return userDetailInfo, nil
}

func (s Service) GetUserByIdOrName(id int, name string) (*User, *entity.ErrCode) {
	user, err := s.Dao.GetUserByIdOrName(id, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s Service) UpdateUserInfoById(user *User) *entity.ErrCode {

	if err := s.Dao.UpdateUserInfoById(user); err != nil {
		return err
	}
	return nil
}

func (s Service) InsertUser(user *User) *entity.ErrCode {
	if err := s.Dao.InsertUser(user); err != nil {
		return err
	}
	return nil
}

//查询绑定人的信息
func (s Service) QueryBindUser(userId int) (*User, *entity.ErrCode) {
	if userId == 0 {
		return nil, errcode.PARAM_ERROR.AddMsg("用户id为空")
	}
	userBind, err := s.Dao.QueryBindUser(userId)
	if err != nil {
		return nil, err
	}
	user, err := s.Dao.GetUserByIdOrName(userBind.BindUserId, "")
	if err != nil {
		return nil, err
	}
	return user, nil
}
