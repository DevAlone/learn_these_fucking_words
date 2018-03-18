import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {RouterModule, Routes} from '@angular/router';
import {StartPageComponent} from './start-page/start-page.component';
import {LoginPageComponent} from './login-page/login-page.component';
import {WordsPageComponent} from './words-page/words-page.component';
import {DebugPageComponent} from "./debug-page/debug-page.component";
import {MyWordsPageComponent} from "./my-words-page/my-words-page.component";
import {RegisterPageComponent} from "./register-page/register-page.component";
import {AuthPageComponent} from "./auth-page/auth-page.component";
import { LearningPageComponent } from "./learning-page/learning-page.component";

const routes: Routes = [
  { path: '', redirectTo: 'index', pathMatch: 'full' },
  { path: 'index', component: StartPageComponent },
  { path: 'login', component: AuthPageComponent },
  { path: 'register', component: AuthPageComponent },
  { path: 'words', component: WordsPageComponent },
  { path: 'my/words', component: MyWordsPageComponent },
  { path: 'learning', component: LearningPageComponent },
  { path: 'debug', component: DebugPageComponent },
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
