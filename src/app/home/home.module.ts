import { FileViewComponent } from './file-view/file-view.component';
import { NoteComponent } from './note/note.component';
import { HomeRouteModule } from './home.routes';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import {PdfViewerComponent} from "ng2-pdf-viewer";
@NgModule({
    imports:[
        HomeRouteModule,
        FormsModule,
        CommonModule,
    ],
    declarations:[
        FileViewComponent
    ],
    providers: [
    ]
})

export class HomeModule{
}