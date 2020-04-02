package photo

import (
	"jaden/we-life/domain/news"
	"jaden/we-life/entity"
	"jaden/we-life/errcode"
)

type IService interface {
	//获取相册列表
	GetDirList(userId int) ([]*Dir, *entity.ErrCode)

	//相册中照片列表
	GetFileList(userId, directoryId int) (*DirResult, *entity.ErrCode)

	//新建相册
	AddDirectory(userId int, dirName string, dirType int) *entity.ErrCode

	//发布动态
	PublishNew(entity *PublishEntity) *entity.ErrCode

	//删除动态
	DeleteNew(newId int) *entity.ErrCode

	//修改相册名称
	UpdateDirName(dirId int, name string) *entity.ErrCode
}

//涉及到首页相册列表
func (s Service) GetDirList(userId int) ([]*Dir, *entity.ErrCode) {
	if userId == 0 {
		return nil, errcode.PARAM_ERROR
	}
	dir, err := s.Dao.GetDirList(userId)
	if err != nil {
		return nil, err
	}
	return dir, nil
}

//相册中的动态列表
func (s Service) GetFileList(userId, directoryId int) (*DirResult, *entity.ErrCode) {

	if userId == 0 || directoryId == 0 {
		return nil, errcode.PARAM_ERROR
	}

	if news, err := s.NewsDao.GetNewsByUserID(userId); err != nil {
		return nil, err
	} else {
		var newIds []int
		var list []List
		for _, info := range news {

			newIds = append(newIds, info.Id)
			n := New{
				NewID:      info.Id,
				Content:    info.Content,
				CreateTime: info.CreateTime,
				Address:    info.Address,
			}
			fileInfo, err := s.Dao.GetFilesByNewId(info.Id)
			//todo 可能news不存在文件
			if err != nil {
				return nil, err
			}
			fls := FileInfoAdapter(fileInfo)
			l := List{
				FileList: fls,
				New:      n,
			}
			list = append(list, l)
		}
		totalFile, err := s.Dao.GetFileCountByDirType(newIds...)
		if err != nil {
			return nil, err
		}

		dirResult := &DirResult{
			Data: Data{
				List:  list,
				Total: totalFile,
			},
		}
		return dirResult, nil

	}
}

func (s Service) AddDirectory(userId int, dirName string, dirType int) *entity.ErrCode {
	if userId == 0 || dirName == "" {
		return errcode.PARAM_ERROR
	}
	if err := s.Dao.InsertDir(dirName, userId, dirType); err != nil {
		return err
	}
	return nil
}

func (s Service) PublishNew(entity *PublishEntity) *entity.ErrCode {
	if entity == nil {
		return errcode.PARAM_ERROR.AddMsg("发布实体信息为空")
	}
	newID, err := s.NewsDao.InsertNew(entity.UserId, entity.Content, entity.Address)
	if err != nil {
		return err
	}
	err = s.Dao.InsertFiles(entity, newID)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteNew(newId int) *entity.ErrCode {
	if newId == 0 {
		return errcode.PARAM_ERROR.AddMsg("动态id异常")
	}
	err := s.NewsDao.DeleteNew(newId)
	if err != nil {
		return err
	}
	err = s.Dao.DeleteFilesByNewId(newId)
	if err != nil {
		return err
	}
	return nil
}
func (s Service) UpdateDirName(dirId int, name string) *entity.ErrCode {
	if dirId == 0 {
		return errcode.PARAM_ERROR.AddMsg("目录id为空")
	}
	err := s.NewsDao.UpdateDirTitle(dirId, name)
	if err != nil {
		return err
	}
	return nil
}

type Service struct {
	Dao     *Dao
	NewsDao *news.Dao
}

func NewService() IService {
	return &Service{
		Dao:     NewDao(),
		NewsDao: news.NewDao(),
	}
}

func FileInfoAdapter(fis []*FileInfo) []*FileList {
	if fis == nil || len(fis) == 0 {
		return nil
	}
	var result []*FileList
	for _, f := range fis {
		fl := &FileList{
			FileID:     f.Id,
			FileName:   f.Name,
			Path:       f.Path,
			UploadTime: int(f.UploadTime),
			ViewCount:  f.ViewCount,
		}
		result = append(result, fl)
	}
	return result
}
