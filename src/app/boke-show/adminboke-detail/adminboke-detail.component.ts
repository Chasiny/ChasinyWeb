import { Data, Code } from './../../app.module';
import { BaseRequestService } from './../../home/services/base-request-service';
import { BoKe } from './../home/struct/data';
import {ActivatedRoute} from "@angular/router";
import { Component, OnInit } from '@angular/core';
@Component({
    selector: 'boke-detail',
    styleUrls: ['./adminboke-detail.component.css'],
    templateUrl: "./adminboke-detail.component.html",
})

export class AdminBoKeDetailComponent implements OnInit {

    getDetailUrl=Data.host+"/api/boke/admindetail"

    constructor(private requestservice:BaseRequestService,private router: ActivatedRoute) {}
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
    id=0;
    title="";
    body="";

    getDetail(){
        this.requestservice.getJSON(this.getDetailUrl+"?bokeid="+this.id,success=>{
            if(success.Code!=Code.SuccessCode){
                alert("系统忙。。。")
                return;
            }
            this.title=success.Data.Title;
            this.body=success.Data.Body;
        },fail=>{
            alert("加载失败！");
        })
    }

}
