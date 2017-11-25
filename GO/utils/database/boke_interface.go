package database

import (
	"errors"
	"../../data"
	"fmt"
	"time"
)

func InsertBoKe(userid int64, title string, body string,status int64) error {
	if (userid < 0) {
		return errors.New("id error")
	}
	if (title == "") {
		return errors.New("title nil")
	}
	if (status < 0) {
		return errors.New("status error")
	}
	sql := "insert into boke(userid,title,body,status) values($1,$2,$3,$4);"
	_, err := Db.Exec(sql, userid, title, body,status)
	return err
}

func UpdateBoKe(userid int64, bokeid int64,title string, body string,status int64) error {
	if (userid < 0) {
		return errors.New("userid error")
	}
	if (bokeid < 0) {
		return errors.New("bokeid error")
	}
	if (title == "") {
		return errors.New("title nil")
	}
	if (status < 0) {
		return errors.New("status error")
	}
	sql := "update boke set title=$1,body=$2,status=$3 where id=$4 and userid =$5;"
	_, err := Db.Exec(sql, title,body,status,bokeid,userid)
	return err
}

func GetBoKeList(userid int64) (bokelist []data.BoKe, err error) {
	var bklist []data.BoKe
	if (userid < 0) {
		return bokelist, errors.New("username nil")
	}
	sql := "select id,title,body,status,createtime from boke where userid = $1"
	rows, err := Db.Query(sql, userid)
	if (err != nil) {
		return bklist, err
	}
	for (rows.Next()) {
		var id,status NullInt
		var createtime time.Time
		var title, body NullStr
		err:=rows.Scan(&id.Nullint64, &title.Nullstr, &body.Nullstr,&status.Nullint64, &createtime)
		if(err!=nil){
			fmt.Println("GetBoKeList error:"+err.Error())
			return bklist,err
		}
		if (!body.Nullstr.Valid) {
			body.Nullstr.String = ""
		}
		bk := data.BoKe{
			BokeId:     id.Nullint64.Int64,
			UserName:   "",
			Title:      title.Nullstr.String,
			Body:       body.Nullstr.String,
			Status:status.Nullint64.Int64,
			CreateTime: createtime.Unix(),
		}
		fmt.Println(bk)
		bklist = append(bklist, bk)
	}
	return bklist, nil
}

func GetBoKe(bokeid int64) (boke data.BoKe, err error) {
	if (bokeid < 0) {
		return boke, errors.New("bokeid nil")
	}
	sql := "select id,title,body,status,createtime from boke where id =$1"
	rows := Db.QueryRow(sql, bokeid)
	var id,status NullInt
	var createtime time.Time
	var title, body NullStr
	err = rows.Scan(&id.Nullint64, &title.Nullstr, &body.Nullstr,&status.Nullint64, &createtime)
	if (err != nil) {
		return boke, err
	}
	if (!body.Nullstr.Valid) {
		body.Nullstr.String = ""
	}
	bk := data.BoKe{
		BokeId:     id.Nullint64.Int64,
		UserName:   "",
		Title:      title.Nullstr.String,
		Body:       body.Nullstr.String,
		CreateTime: createtime.Unix(),
	}
	return bk, nil
}

func DeleteBoKe(userid int64,bokeid int64) error {
	if (userid < 0) {
		return errors.New("id error")
	}
	if (bokeid < 0) {
		return errors.New("bokeid error")
	}
	sql := "delete from boke where id = $1 and userid = $2;"
	_, err := Db.Exec(sql, bokeid, userid)
	return err
}