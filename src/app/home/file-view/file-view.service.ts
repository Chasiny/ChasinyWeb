import { BaseRequestService } from './../services/base-request-service';
import { Data } from './../../app.module';
import { isUndefined } from 'util';
import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Headers, RequestOptions } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

@Injectable()
export class FileViewService extends BaseRequestService {
    private host = Data.host;

    constructor() {super(); }

    getFilesDirs(dir:string,success:Function,fail:Function) {
        let url:string;
        if(dir==""){
            url=this.host+"/getlist"
        }else{
            url=this.host+"/getlist?dir="+dir
        }
        this.getJSON(url,data=>{
            success(data);
        },err=>{
            fail(err);
        });
    }

    private extractData(res: Response) {
        let body = res.json();
        return body || {};
    }

    private handleError(error: Response | any) {
        let errMsg: string;
        if (error instanceof Response) {
            const body = error.json() || '';
            const err = body.error || JSON.stringify(body);
            errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
        } else {
            errMsg = error.message ? error.message : error.toString();
        }
        console.error(errMsg);
        return Observable.throw(errMsg);
    }

    public Download(name:string){
        let url=this.host+"/download?filename="+name;
        /*
        const a: HTMLAnchorElement = document.createElement('a');
        a.href = url;
        a.download = 'download'
        a.click();
        a.remove();
        console.log("download:" + a.href);
        */
        window.open(url);
    }
}


