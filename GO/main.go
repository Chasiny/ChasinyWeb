package main

import (
	"log"
	"net/http"
	"strconv"
	"./controller/filecontroller"
	"fmt"
	"./view"
	mux "github.com/gorilla/mux"
	"./utils"
	"./data"
	"./utils/database"
	"./controller/login"
	"./controller/dbfile"
	"./controller/session"
	"./controller/boke"
)


func main() {
	view.Init()
	database.Init()
	session.Init()
	r := mux.NewRouter()
	r.PathPrefix("/home").HandlerFunc(view.LoadTemplate)
	r.PathPrefix("/login").HandlerFunc(view.LoadTemplate)
	r.PathPrefix("/boke").HandlerFunc(view.LoadTemplate)
	InitLoginFunc(r)
	InitDBFileFunc(r)
	InitBoKeFunc(r)
	InitFileFunc(r)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(utils.GetCurrentDirectory()+"/dist"))))
	fmt.Println("start...")
	err := http.ListenAndServe(":"+strconv.Itoa(data.Port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func InitLoginFunc(r *mux.Router ){
	r.HandleFunc("/api/login", login.Login)
}

func InitDBFileFunc(r *mux.Router ){
	r.HandleFunc("/api/dbfile/download", dbfile.DownloadDBFile())
	r.HandleFunc("/api/dbfile/upload", dbfile.UploadDBFile())
}


func InitBoKeFunc(r *mux.Router ){
	r.HandleFunc("/api/boke/adminbokelist", boke.GetAdminBokeList)
	r.HandleFunc("/api/boke/admindetail", boke.GetAdminBokeDetail)
	r.HandleFunc("/api/boke/list", boke.GetBokeList())
	r.HandleFunc("/api/boke/detail", boke.GetAdminBokeDetail)
	r.HandleFunc("/api/boke/add", boke.AddBoke())
	r.HandleFunc("/api/boke/edit", boke.EditBoke())
	r.HandleFunc("/api/boke/delete", boke.DeleteBoke())
}


func InitFileFunc(r *mux.Router ){
	r.HandleFunc("/test", filecontroller.Test())
	r.HandleFunc("/getlist", filecontroller.GetFileList())
	r.HandleFunc("/download", filecontroller.DownloadFile())
	r.HandleFunc("/upload", filecontroller.UploadFile())
}