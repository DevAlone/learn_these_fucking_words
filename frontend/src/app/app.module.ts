import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { StartPageComponent } from './start-page/start-page.component';
import { MainMenuComponent } from './main-menu/main-menu.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { WordsPageComponent } from './words-page/words-page.component';
import { HttpClientModule } from '@angular/common/http';
import { JwtModule } from '@auth0/angular-jwt';
import { FormsModule } from '@angular/forms';
import { ApiService } from './api.service';
import { MessagesComponent } from './messages/messages.component';
import { MessageService } from "./message.service";
import { DebugPageComponent } from './debug-page/debug-page.component';
import { MyWordsPageComponent } from './my-words-page/my-words-page.component';
import { RegisterPageComponent } from './register-page/register-page.component';
import { AuthPageComponent } from './auth-page/auth-page.component';
import { LearningPageComponent } from './learning-page/learning-page.component';
import { WordInfoComponent } from './word-info/word-info.component';
import { WordInfoPearsonComComponent } from './word-info-pearson-com/word-info-pearson-com.component';
import { MaterialModule } from './material.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AudioPlayButtonComponent } from './audio-play-button/audio-play-button.component';
import { WordInfoPearsonComExampleComponent } from './word-info-pearson-com-example/word-info-pearson-com-example.component';

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
    AuthPageComponent,
    LearningPageComponent,
    WordInfoComponent,
    WordInfoPearsonComComponent,
    AudioPlayButtonComponent,
    WordInfoPearsonComExampleComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: function() { return localStorage.getItem('access_token'); },
        whitelistedDomains: ['localhost:8080'],
        blacklistedRoutes: ['localhost:8080/login'],
        // throwNoTokenError: true,
      }
    }),
    MaterialModule,
    BrowserAnimationsModule
  ],
  providers: [
    ApiService,
    MessageService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
