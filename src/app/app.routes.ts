import { AdminBoKeDetailComponent } from './boke-show/adminboke-detail/adminboke-detail.component';
import { AdminBoKeShowComponent } from './boke-show/adminboke-show.component';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { Routes } from '@angular/router';
import { DataResolver } from './app.resolver';
import{LoginComponent}from './login/login.component'

export const appRoutes: Routes = [
    { path: '',
      redirectTo: '/login',
      pathMatch: 'full',
    },
    { path: 'home', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'boke', component: AdminBoKeShowComponent },
    { path: 'bokedetail', component: AdminBoKeDetailComponent },
  ];