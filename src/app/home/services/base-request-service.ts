
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/toPromise';

// 向服务器请求数据的基类服务(尽量不要直接使用或修改它，继承并使用它！！)
@Injectable()
export class BaseRequestService {
  // 请求最初始方法
  // sentences_url:链接，method：请求方法，data：发送数据，success：成功回调函数，fail：失败回调函数，
  // responseType：回复内容形式，sync：是否同步请求,uploadProgress 回调进度,downLoadProgress 下载回调进度
  baseRequest(url: string, method: string , data?: FormData | string, success?: Function, fail?: Function, state?: Function,
              responseType?: string, uploadProgress?:
                Function, downloadProgress?: Function, sync: boolean= true, header?: [{name: string, value: string}]): void {
    const xhr = new XMLHttpRequest();
    if (responseType) {
      xhr.responseType = responseType;
    } else {
      xhr.responseType =  '';
    }
    xhr.withCredentials=true;
    xhr.open(method, url, sync);
    if (header) {
      for (const h of header){
        xhr.setRequestHeader(h.name, h.value);
      }
    }
    xhr.send(data);
    if (!success && !fail) {
      return;
    }
    if (uploadProgress) {
      xhr.upload.addEventListener('progress', function (ev: ProgressEvent) {
        uploadProgress(ev.loaded, ev.total);
      }, false);
      // xhr.upload.onprogress = function (ev: ProgressEvent) {
      //   uploadProgress(ev.loaded, ev.total);
      // };
    }
    if (downloadProgress) {
      xhr.addEventListener('progress', function (evn) {
        downloadProgress(evn.loaded, evn.total);
      });
    }
    xhr.onreadystatechange = function () {
      if (!(this.readyState === 4 && this.status === 200)) {
        if (state && this.status === 200) {
          state(this.readyState);
        }
        if (fail && this.status !== 200) {
          fail(this.status);
        }
      } else {
        if (success) {
          success(this.response);
        }
      }
    };
  }
  // GET 请求
  // sentences_url:链接，success：成功回调函数，fail：失败回调函数，
  // responseType：回复内容形式，sync：是否同步请求
  get(url: string, success?: Function, fail?: Function, state?: Function,
      responseType?: string, sync: boolean= true, header?: [{name: string, value: string}]) {
    this.baseRequest(url, 'get' , null, success, fail, state, responseType, null, null, sync, header);
  }
  // POST 请求
  // sentences_url:链接，data：发送数据，success：成功回调函数，fail：失败回调函数，
  // responseType：回复内容形式，sync：是否同步请求
  post(url: string, data?: FormData | string, success?: Function, fail?: Function, state?: Function,
       responseType?: string, sync: boolean= true, header?: [{name: string, value: string}]) {
    this.baseRequest(url, 'post', data, success, fail, state, responseType, null, null, sync , header);
  }
  // POST 请求 获取JSON
  postJSON(url: string, data?: FormData | any, success?: Function, fail?: Function, sync: boolean= true,
           header?: [{name: string, value: string}]) {
    this.baseRequest(url, 'post', data, success, fail, null, 'json', null, null, sync , header);
  }
  //Get请求 获取JSON
  getJSON(url: string, success?: Function, fail?: Function, sync: boolean= true,
          header?: [{name: string, value: string}]){
    this.baseRequest(url, 'get',null, success, fail, null, 'json', null, null, sync , header);
  }
}
//XMLHttpRequestResponseType
