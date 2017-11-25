import { BoKe } from './../home/struct/data';
import { BaseRequestService } from './../home/services/base-request-service';
import { Data, Code } from './../app.module';
import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
@Component({
    selector: 'boke-show',
    styleUrls: ['./adminboke-show.component.css'],
    templateUrl: "./adminboke-show.component.html",
})

export class AdminBoKeShowComponent implements OnInit {

    getListUrl=Data.host+"/api/boke/adminbokelist";
    list=[];

    constructor(private requestservice:BaseRequestService,private router: Router) {}
    public ngOnInit() {this.GetList();}

    OpenBoKe(id:number){
        console.log(id);
        this.router.navigate(["./bokedetail"],{queryParams:{bokeid:id}});
    }

    GetList(){
        this.requestservice.getJSON(this.getListUrl,success=>{
            if(success.Code!=Code.SuccessCode){
                alert("系统忙。。。")
                return;
            }
            console.log("load success");
            this.list=success.Data;
            console.log(this.list);
        },fail=>{
            alert("加载失败！");
        });
    }
}
