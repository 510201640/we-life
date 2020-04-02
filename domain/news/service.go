package news

import (
	"jaden/we-life/domain/user"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
)

type IService interface {
	//根据用户id查询动态(与绑定人的所有动态)
	GetNewsByUserID(userId int) ([]*News, *entity.ErrCode)

	//新增一条动态
	InsertNew(userId int, content string, address string) (int, *entity.ErrCode)
}

func (s Service) InsertNew(userId int, content string, address string) (int, *entity.ErrCode) {
	if userId == 0 {
		return 0, errcode.PARAM_ERROR
	}
	id, err := s.Dao.InsertNew(userId, content, address)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//获取 userId 与 绑定人的news
func (s Service) GetNewsByUserID(userId int) ([]*News, *entity.ErrCode) {
	bindUser, err := s.UserService.QueryBindUser(userId)
	if err != nil {
		return nil, err
	}
	news, err := s.Dao.GetNewsByUserIDs([]int{userId, bindUser.Id})
	if err != nil {
		return nil, err
	}
	return news, nil
}

type Service struct {
	Dao         *Dao
	UserService user.IService
}

func NewService() IService {
	return &Service{
		Dao:         NewDao(),
		UserService: user.NewService(),
	}
}
