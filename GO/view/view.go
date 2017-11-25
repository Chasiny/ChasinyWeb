package view

import (
	"github.com/golang/glog"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"../utils"
	//"../../session"
	"fmt"
)

var (
	indexTemplate *template.Template
	templates     *template.Template
)

func Init() {
	var allfile []string
	files, err := ioutil.ReadDir(utils.GetCurrentDirectory()+"/dist/")
	if err != nil {
		glog.Errorln(err)
		return
	}
	for _, file := range files {
		fileName := file.Name()
		fmt.Println(utils.GetCurrentDirectory()+"/dist/"+fileName)
		if strings.HasSuffix(fileName, ".html") {
			allfile = append(allfile, utils.GetCurrentDirectory()+"/dist/"+fileName)
		}
	}
	templates = template.Must(template.ParseFiles(allfile...))
	indexTemplate = templates.Lookup("index.html")
}

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	/*
		if err:=session.JudgeLegality(r);err!=nil{
			http.NotFound(w,r)
			return
		}
	*/
	if r.Method == "GET" {
		indexTemplate.Execute(w, nil)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
