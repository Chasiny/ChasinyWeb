import { BaseRequestService } from './../home/services/base-request-service';
import { Data, Code } from './../app.module';
import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
@Component({
    selector: 'login',
    styleUrls: ['./login.component.css'],
    templateUrl: "./login.component.html",
})

export class LoginComponent implements OnInit {

    constructor(private requestservice:BaseRequestService,private router: Router) {}
    public ngOnInit() {}

    loginurl=Data.host+"/api/login";

    name="";
    pwd="";

    Login(){
        if(this.name===""||this.pwd===""){
            console.log("nameor pwd nil");
            return;
        }
        let fd=new FormData;
        fd.append("name",this.name);
        fd.append("pwd",this.pwd);
        this.requestservice.postJSON(this.loginurl,fd,success=>{
            if(success.Code===Code.SuccessCode){
                this.router.navigate(['home']);
            }else{
                console.log("pwd error");
            }
        },fail=>{
            console.log("error");
        })
    }

}
