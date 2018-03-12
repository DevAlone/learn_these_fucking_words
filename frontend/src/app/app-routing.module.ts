import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {RouterModule, Routes} from '@angular/router';
import {StartPageComponent} from './start-page/start-page.component';
import {LoginPageComponent} from './login-page/login-page.component';
import {WordsPageComponent} from './words-page/words-page.component';

const routes: Routes = [
  { path: '', redirectTo: 'index', pathMatch: 'full' },
  { path: 'index', component: StartPageComponent },
  { path: 'login', component: LoginPageComponent },
  { path: 'words', component: WordsPageComponent },
  // { path: 'channel/:id', component: ChannelDetailComponent },
];

@NgModule({
  // imports: [
  //   // CommonModule
  //   RouterModule.forRoot(routes, { useHash: true })
  // ],
  // declarations: []
  exports: [ RouterModule ],
  imports: [ RouterModule.forRoot(routes, { useHash: true }) ],
})
export class AppRoutingModule { }
