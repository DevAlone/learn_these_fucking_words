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
import {AuthService} from './auth.service';
import {FormsModule} from '@angular/forms';
import {ApiService} from './api.service';


export function tokenGetter() {

}


@NgModule({
  declarations: [
    AppComponent,
    StartPageComponent,
    MainMenuComponent,
    LoginPageComponent,
    WordsPageComponent
  ],
  imports: [
    FormsModule,
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: () => localStorage.getItem('access_token'),
        whitelistedDomains: ['localhost:8080'],
        blacklistedRoutes: ['localhost:8080/login'],
        // throwNoTokenError: true,
      }
    })
  ],
  providers: [
    AuthService,
    ApiService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
