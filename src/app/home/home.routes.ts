import { BoKeDetailComponent } from './boke/detail/boke-detail.component';
import { BoKeListComponent } from './boke/list/boke-list.component';
import { BoKeAddComponent } from './boke/add/boke-add.component';
import { HomeComponent } from './home.component';
import { FileViewComponent } from './file-view/file-view.component';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

@NgModule({
    imports: [
        RouterModule.forRoot([
            {
            path: 'home',
            component: HomeComponent,
            children: [
                {
                    path: 'boke/add',
                    component: BoKeAddComponent,
                },
                {
                    path: 'boke/list',
                    component: BoKeListComponent,
                },
                {
                    path: 'boke/detail',
                    component: BoKeDetailComponent,
                },
                {
                    path: 'fileview',
                    component: FileViewComponent,
                }
            ]
            }
        ]),
    ],
    exports: [
        RouterModule,
    ]
})

export class HomeRouteModule {

}