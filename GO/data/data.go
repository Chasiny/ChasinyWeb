package data

type FileList struct {
	FileName []string `json:"FileName"`
	FileDir  []string `json:"FileDir"`
	CurDir   string   `json:"CurDir"`
}

type DBFileList struct {
	FileId   int64  `json:"FileId"`
	FileName string `json:"FileName"`
}

type BoKe struct {
	BokeId int64  `json:"BokeId"`
	UserName string  `json:"UserName"`
	Title string  `json:"Title"`
	Body string  `json:"Body"`
	Status int64  `json:"Status"`
	CreateTime int64 `json:"CreateTime"`
}

const (
	Host    = "http://127.0.0.1"
	Port    = 8088
	WebHost = "http://127.0.0.1:3000"
	//WebHost="http://119.29.91.244:8088"
	DbHost = "127.0.0.1"
	DbPort = 5432
	DbUser = "postgres"
	DbPwd  = "caihongye"
	DbName = "chy"
)
