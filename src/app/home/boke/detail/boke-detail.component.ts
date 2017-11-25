import { Data, Code } from './../../../app.module';
import { BaseRequestService } from './../../services/base-request-service';
import {ActivatedRoute,Router} from "@angular/router";
import { Component, OnInit } from '@angular/core';
@Component({
    selector: 'boke-detail',
    styleUrls: ['./boke-detail.component.css'],
    templateUrl: "./boke-detail.component.html",
})

export class BoKeDetailComponent implements OnInit {

    getDetailUrl=Data.host+"/api/boke/detail"
    deleteBokeUrl=Data.host+"/api/boke/delete"
    id=0;
    title="";
    body="";

    constructor(private requestservice:BaseRequestService,private router: ActivatedRoute,private navRouter:Router) {}
    public ngOnInit() {
        this.router.queryParams.subscribe((res)=>{
            if(res['bokeid']==''||res['bokeid']==undefined){
                alert("参数错误！");
                return;
            }
            this.id=res['bokeid'];
            this.getDetail();
        });
    }

    getDetail(){
        this.requestservice.getJSON(this.getDetailUrl+"?bokeid="+this.id,success=>{
            if(success.Code!=Code.SuccessCode){
                alert("系统忙。。。")
                return;
            }
            this.title=success.Data.Title;
            this.body=success.Data.Body;
        },fail=>{
            alert("网络连接失败！");
        })
    }

    Edit(){
        this.navRouter.navigate(["./home/boke/add"],{queryParams:{type:"edit",bokeid:this.id}});
    }

    Delete(){
        this.requestservice.getJSON(this.deleteBokeUrl+"?bokeid="+this.id,success=>{
            if(success.Code!=Code.SuccessCode){
                alert("系统忙。。。")
                return;
            }
            alert("删除成功!");
            this.navRouter.navigate(["./home/boke/list"]);
        },fail=>{
            alert("网络连接失败！");
        })
    }

}
