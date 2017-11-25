import { BoKeDetailComponent } from './home/boke/detail/boke-detail.component';
import { AdminBoKeShowComponent } from './boke-show/adminboke-show.component';
import { AdminBoKeDetailComponent } from './boke-show/adminboke-detail/adminboke-detail.component';
import { BoKeListComponent } from './home/boke/list/boke-list.component';
import { BoKeAddComponent } from './home/boke/add/boke-add.component';

import { BaseRequestService } from './home/services/base-request-service';
import { FileViewService } from './home/file-view/file-view.service';
import { HomeModule } from './home/home.module';
import { HomeComponent } from './home/home.component';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import {NgModule,ApplicationRef} from '@angular/core';
import {removeNgStyles,createNewHosts,createInputTransfer} from '@angularclass/hmr';
import {RouterModule,PreloadAllModules} from '@angular/router';
import { ENV_PROVIDERS } from './environment';
import { appRoutes } from './app.routes';
import { AppComponent } from './app.component';
import { APP_RESOLVER_PROVIDERS } from './app.resolver';
import { AppState, InternalStateType } from './app.service';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import{LoginComponent}from './login/login.component'
import '../styles/styles.scss';
import '../styles/headings.css';
import { MarkdownModule } from 'angular2-markdown';


const APP_PROVIDERS = [
  ...APP_RESOLVER_PROVIDERS,
  AppState
];

type StoreType = {
  state: InternalStateType,
  restoreInputValues: () => void,
  disposeOldHosts: () => void
};


@NgModule({
  bootstrap: [ AppComponent ],
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    AdminBoKeDetailComponent,
    AdminBoKeShowComponent,
    BoKeListComponent,
    BoKeAddComponent,
    BoKeDetailComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HomeModule,
    HttpModule,
    BrowserAnimationsModule,
    MarkdownModule.forRoot(),
    RouterModule.forRoot(appRoutes),
  ],
  providers: [
    BaseRequestService,
    FileViewService,
    ENV_PROVIDERS,
    APP_PROVIDERS
  ]
})

export class AppModule {

  constructor(
    public appRef: ApplicationRef,
    public appState: AppState
  ) {}

  public hmrOnInit(store: StoreType) {
    if (!store || !store.state) {
      return;
    }
    console.log('HMR store', JSON.stringify(store, null, 2));
    this.appState._state = store.state;
    if ('restoreInputValues' in store) {
      let restoreInputValues = store.restoreInputValues;
      setTimeout(restoreInputValues);
    }

    this.appRef.tick();
    delete store.state;
    delete store.restoreInputValues;
  }

  public hmrOnDestroy(store: StoreType) {
    const cmpLocation = this.appRef.components.map((cmp) => cmp.location.nativeElement);
    const state = this.appState._state;
    store.state = state;
    store.disposeOldHosts = createNewHosts(cmpLocation);
    store.restoreInputValues  = createInputTransfer();
    removeNgStyles();
  }

  public hmrAfterDestroy(store: StoreType) {
    store.disposeOldHosts();
    delete store.disposeOldHosts;
  }

}

export class Data{
  //public static host="http://127.0.0.1:8088";
  public static host="http://chasiny.top:8088";
}

export class Code{
  public static SuccessCode=10000;
	public static ParmErr=10001;
	public static PwdErr=10002;
}