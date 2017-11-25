import { Router } from '@angular/router';
import { BaseRequestService } from './../../services/base-request-service';
import { Data, Code } from './../../../app.module';
import { BoKe } from './../../struct/data';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
@Component({
    encapsulation: ViewEncapsulation.None,
    selector: 'boke-list',
    styleUrls: ['./boke-list.component.css','./../detail/boke-detail.component.css'],
    templateUrl: "./boke-list.component.html",
})

export class BoKeListComponent implements OnInit {
    
    getListUrl=Data.host+"/api/boke/list";
    list=[];

    constructor(private requestservice:BaseRequestService,private router: Router) {}
    public ngOnInit() {this.GetList();}

    OpenBoKe(id:number){
        console.log(id);
        this.router.navigate(["./home/boke/detail"],{queryParams:{bokeid:id}});
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

    AddBoke(){
        this.router.navigate(["./home/boke/add"],{queryParams:{type:"add",bokeid:0}});
    }
}
