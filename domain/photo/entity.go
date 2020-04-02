package photo

type Dir struct {
	Id            int
	DirectionName string `json:"directionName"`
	DirFileCount  int    `json:"dirFileCount"`
	CreateTime    int    `json:"createTime"`
	UserId        int    `json:"userId"`
	IsDelete      int
	DirType       int `json:"dirType"`
}

type DirResult struct {
	Data Data `json:"data"`
}
type FileList struct {
	FileID     int    `json:"fileId"`
	FileName   string `json:"fileName"`
	Path       string `json:"path"`
	UploadTime int    `json:"uploadTime"`
	ViewCount  int    `json:"view_count"`
}
type New struct {
	NewID      int    `json:"newId"`
	Content    string `json:"content"`
	CreateTime int    `json:"createTime"`
	Address    string `json:"address"`
}
type List struct {
	FileList []*FileList `json:"fileList"`
	New      New         `json:"new"`
}
type Data struct {
	List  []List `json:"list"`
	Total int    `json:"total"`
}

type FileInfo struct {
	Id          int
	Name        string
	Path        string
	ViewCount   int
	UploadTime  int64
	IsDelete    int
	DirectoryId int
	UserId      int
	NewId       int
}

type PublishEntity struct {
	UserId           int
	DirId            int
	Content          string
	Address          string
	UploadFileEntity []UploadFileEntity
}

type UploadFileEntity struct {
	FileName string
	FilePath string
}
