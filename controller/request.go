package controller

type NewRequest struct {
	UserId      int    `form:"userId"`
	DirectoryId int    `form:"directoryId"`
	Session     string `form:"session"`
}

type AddDirRequest struct {
	DirName string `json:"dirName"`
	UserID  int    `json:"userId"`
}

type DeleteNewReqest struct {
	UserId int `json:"userId"`
	NewId  int `json:"newId"`
}

type UpdateDirNameRequest struct {
	DirId    int    `json:"dirId"`
	NewTitle string `json:"newTitle"`
}

type UserLoginRequest struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
}

type UserBindRequest struct {
	UserId     int `json:"userId"`
	BindUserId int `json:"bindUserId"`
}
