package dbfile

import (
	"net/http"
	"fmt"
	"../../controller"
)

func DownloadDBFile() http.HandlerFunc {
	return controller.GateWay(downloadDBFile)
}
func downloadDBFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download")
}

func UploadDBFile() http.HandlerFunc {
	return controller.GateWay(uploadDBFile)
}
func uploadDBFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download")
}