package filecontroller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"../../data"
	"../../utils"
	"encoding/json"
	//"../../errorcode"
	ioutil "io/ioutil"
	"../../controller"
)

func Test() http.HandlerFunc {
	return controller.GateWay(test)
}
func test(w http.ResponseWriter, r *http.Request) {
	files, _, _ := utils.ListDir(utils.GetCurrentDirectory())
	fmt.Println("Test")
	f, err := os.Open(files[1])
	if err != nil {
		fmt.Println("error file :" + files[1])
		return
	}
	data := bufio.NewReader(f)

	data.WriteTo(w)
}

func GetFileList() http.HandlerFunc {
	return controller.GateWay(getFileList)
}
func getFileList(w http.ResponseWriter, r *http.Request) {

	dir := r.FormValue("dir")
	if dir == "" {
		dir = utils.GetCurrentDirectory()
	}
	files, filedirs, _ := utils.ListDir(dir)
	postdata := data.FileList{
		FileName: files,
		FileDir:  filedirs,
		CurDir:   dir,
	}
	databuf, err := json.Marshal(postdata)
	if err != nil {
		fmt.Print(err)
		return
	}
	w.Write(databuf)
}

func DownloadFile() http.HandlerFunc {
	return controller.GateWay(downloadFile)
}
func downloadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download")
	filename := r.FormValue("filename")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error file :" + filename)
		w.Write([]byte("error"))
		return
	}
	data := bufio.NewReader(f)
	w.Header().Set("Content-Disposition", "attachment; filename="+utils.GetNameFromDir(filename))
	data.WriteTo(w)
}

func UploadFile() http.HandlerFunc {
	return controller.GateWay(uploadFile)
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("dir")
	fmt.Println("UploadFile:" + dir)
	file, fileheader, err := r.FormFile("file")
	if (err != nil) {
		fmt.Println("r.FormFile error:", err.Error())
	}
	filename := dir + "/" + fileheader.Filename
	var f *os.File
	if checkFileIsExist(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Println("文件存在");
	} else {
		f, err = os.Create(filename)
		fmt.Println("文件不存在");
	}
	if(err!=nil){
		fmt.Println("openfile error",err.Error())
		return
	}
	writer := bufio.NewWriter(f)
	buf, err := ioutil.ReadAll(file)
	if(err!=nil){
		fmt.Println("ReadAll error",err.Error())
		return
	}
	n4, err := writer.Write(buf)
	if(err!=nil){
		fmt.Println("writer.Write(buf) error",err.Error())
		return
	}
	fmt.Println("写入 %d 个字节n", n4)
	writer.Flush()
	f.Close()
	w.Write([]byte("ok"))
}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

