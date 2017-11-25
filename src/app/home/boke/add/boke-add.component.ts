import { Router } from '@angular/router';
import { ActivatedRoute } from '@angular/router';
import { BaseRequestService } from './../../services/base-request-service';

import { Data, Code } from './../../../app.module';
import { BoKe } from './../../struct/data';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
@Component({
    encapsulation: ViewEncapsulation.None,
    selector: 'boke-add',
    styleUrls: ['./boke-add.component.css'],
    templateUrl: "./boke-add.component.html",
})

export class BoKeAddComponent implements OnInit {
    addUrl = Data.host + "/api/boke/add"
    getDetailUrl=Data.host+"/api/boke/detail"
    editUrl=Data.host+"/api/boke/edit"
    public textData = "";
    public title = "";
    type="";
    bokeid=0;

    constructor(private baseservice: BaseRequestService,private router: ActivatedRoute,private navRouter:Router) { }
    public ngOnInit() {
        this.router.queryParams.subscribe((res)=>{
            if(res['type']==''||res['type']==undefined){
                alert("参数错误！");
                return;
            }
            this.type=res['type'];
            this.bokeid=res['bokeid'];
            if(this.type=="edit"){
                console.log(this.bokeid);
                this.getDetail();
            }
        });
    }

    AddBoke() {
        if (this.textData == ""||this.title=="") {
            console.log("textdata or title nil");
            return
        }
        let fd =new FormData();
        fd.append("title",this.title);
        fd.append("body",this.textData);
        fd.append("status","1");
        this.baseservice.postJSON(this.addUrl,fd , success => {
            console.log(success.Code);
            if(success.Code==Code.SuccessCode){
                alert("添加成功！");
                this.navRouter.navigate(["./home/boke/list"]);
                return;
            }
            alert("添加失败！");
            return;
        }, fail => {
            console.log(fail);
            alert("网络连接失败！");
            return;
        })
    }

    SaveBoke(){
        if (this.textData == ""||this.title=="") {
            console.log("textdata or title nil");
            return
        }
        let fd =new FormData();
        fd.append("bokeid",this.bokeid.toString());
        fd.append("title",this.title);
        fd.append("body",this.textData);
        fd.append("status","1");
        this.baseservice.postJSON(this.editUrl,fd , success => {
            console.log(success.Code);
            if(success.Code==Code.SuccessCode){
                alert("保存成功！");
                this.navRouter.navigate(["./home/boke/detail"],{queryParams:{bokeid:this.bokeid}});
                return;
            }
            alert("保存失败！");
            return;
        }, fail => {
            console.log(fail);
            alert("网络连接失败！");
            return;
        })
    }

    getDetail(){
        this.baseservice.getJSON(this.getDetailUrl+"?bokeid="+this.bokeid,success=>{
            if(success.Code!=Code.SuccessCode){
                alert("系统忙。。。")
                return;
            }
            this.title=success.Data.Title;
            this.textData=success.Data.Body;
        },fail=>{
            alert("网络连接失败！");
        })
    }

}
