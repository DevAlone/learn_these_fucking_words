import { Injectable } from '@angular/core';
import { Router } from '@angular/router'
// import {Observable} from 'rxjs/Observable';
import {HttpClient} from '@angular/common/http';
import {catchError} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';
import { Observable } from "rxjs/Rx"
import {MessageService} from "./message.service";

@Injectable()
export class ApiService {
  public authorized: boolean = false;

  constructor(
    private http: HttpClient,
    private messageService: MessageService,
    private router: Router,
  ) {
    if (localStorage.getItem('access_token'))
      this.authorized = true;
  }

  public login(username: string, password: string): Observable<void> {
    return this.post('login', {username: username, password: password}).map(res => {
      if (res.hasOwnProperty('token')) {
        localStorage.setItem('access_token', res.token);
        this.authorized = true;
      }
    });
  }

  public logout() {
    localStorage.removeItem('access_token');
    this.authorized = false;
  }

  public register(username: string, password: string): Observable<void> {
    return this.post('register', {username: username, password: password}).map(res => {
      if (res.hasOwnProperty('token')) {
        localStorage.setItem('access_token', res.token);
        this.authorized = true;
      }
    });
  }

  public get(url: string): Observable<any> {
    if (!url.startsWith('/')) {
      url = '/' + url;
    }

    url = 'http://localhost:8080/api/v1' + url;

    // return this.http.get(url, ApiConfig.HTTP_OPTIONS).pipe(
    return this.http.get(url).pipe(
      // tap(_ => console.log("")),
      catchError(this.handleError<any[]>('get', []))
    );
  }

  public post(url: string, data: any): Observable<any> {
    if (!url.startsWith('/')) {
      url = '/' + url;
    }

    url = 'http://localhost:8080/api/v1' + url;

    return this.http.post(url, data).catch(error => {
      this.handleError<any[]>('post', []);
      return Observable.throw(error);
    });

    // this.messageService.error("aaaaaaaaa");
    // return this.http.post(url, data).pipe(
    //   // tap(_ => console.log("")),
    //   catchError(this.handleError<any[]>('post', []))
    // );
  }

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      if (error.status == 401) {
        localStorage.removeItem('access_token');
        this.authorized = false;
        this.router.navigateByUrl('/login');
      }

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      // this.log(`${operation} failed: ${error.message}`);

      this.messageService.error(error.error.error_message);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
