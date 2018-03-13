import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { StartPageComponent } from './start-page/start-page.component';
import { MainMenuComponent } from './main-menu/main-menu.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { WordsPageComponent } from './words-page/words-page.component';
import {HttpClientModule} from '@angular/common/http';
import {JwtModule} from '@auth0/angular-jwt';
import {FormsModule} from '@angular/forms';
import {ApiService} from './api.service';
import { MessagesComponent } from './messages/messages.component';
import {MessageService} from "./message.service";
import { DebugPageComponent } from './debug-page/debug-page.component';
import { MyWordsPageComponent } from './my-words-page/my-words-page.component';
import { RegisterPageComponent } from './register-page/register-page.component';
import { AuthPageComponent } from './auth-page/auth-page.component';


@NgModule({
  declarations: [
    AppComponent,
    StartPageComponent,
    MainMenuComponent,
    LoginPageComponent,
    WordsPageComponent,
    MessagesComponent,
    DebugPageComponent,
    MyWordsPageComponent,
    RegisterPageComponent,
    AuthPageComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: () => localStorage.getItem('access_token'),
        whitelistedDomains: ['localhost:8080'],
        blacklistedRoutes: ['localhost:8080/login'],
        // throwNoTokenError: true,
      }
    }),
  ],
  providers: [
    ApiService,
    MessageService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
