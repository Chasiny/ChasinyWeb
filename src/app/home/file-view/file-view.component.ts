import { isUndefined } from 'util';
import { AppModule, Data } from './../../app.module';

import { FileData } from './file-data';
import { FileViewService } from './file-view.service';
import { Component, OnInit } from '@angular/core';
import { MdButtonModule, MdCheckboxModule } from '@angular/material-ui';
@Component({
    selector: 'file-view',
    styleUrls: ['./../../app.component.css'],
    templateUrl: "./file-view.component.html",
})

export class FileViewComponent implements OnInit {

    fileData: FileData = { Dir: [], DirName: [], File: [], FileName: [] };
    curDir = "";
    uploadFileUrl = Data.host + "/upload";

    constructor(private fileservice: FileViewService) {
        this.getFilesDirs("");
    }
    public ngOnInit() { }

    download(name: string) {
        console.log("download file:", name);
        if (name == "") {
            console.log("name nil");
            return;
        }
        this.fileservice.Download(name);
    }

    getFilesDirs(dir: string) {
        console.log("getFilesDirs", dir);
        this.fileservice.getFilesDirs(dir, success => {
            this.fileData.File = success.FileName;
            this.fileData.Dir = success.FileDir;
            this.curDir = success.CurDir;
            this.initFileDirName();
            console.log(this.fileData);
        }, fail => {
            console.log("getFilesDirs fail", fail);
        });
    }

    getLastDirs() {
        if (this.curDir == "") {
            console.log("curdir nil");
            return;
        }
        let dir = "";
        let spitdir = this.curDir.split("/", -1);
        for (let i = 1; i < spitdir.length - 1; i++) {
            dir = dir + "/" + spitdir[i];
        }
        console.log("fatherdir:", dir);
        this.getFilesDirs(dir);
    }

    getFileName(dir: string): string {
        if (dir == "") {
            console.log("curdirdir nil");
            return;
        }
        let spitdir = this.curDir.split("/", -1);
        return spitdir[spitdir.length - 1];
    }

    initFileDirName() {
        this.fileData.DirName = [];
        this.fileData.FileName = [];
        try {
            for (let i = 0; i < this.fileData.Dir.length; i++) {
                let spitDir = this.fileData.Dir[i].split("/", -1);
                this.fileData.DirName.push(spitDir[spitDir.length - 1]);
            }
        } catch (error) {
        }

        try {
            for (let i = 0; i < this.fileData.File.length; i++) {
                let spitDir = this.fileData.File[i].split("/", -1);
                this.fileData.FileName.push(spitDir[spitDir.length - 1]);
            }
        } catch (error) {
        }
    }

    uploadFile($event: any): void {
        let formData=new FormData;
        console.log("uploadfile:",$event.target.files[0].name);
        formData.append('file', $event.target.files[0], $event.target.files[0].name);
        formData.append('dir', this.curDir);
        this.fileservice.baseRequest(this.uploadFileUrl, 'post', formData, (resp) => {
            this.getFilesDirs(this.curDir);
        }, (code) => { 
        }, null, 'json');
    }
}
