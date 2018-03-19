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
  public authorized = false;

  constructor(
    private http: HttpClient,
    private messageService: MessageService,
    private router: Router,
  ) {
    if (localStorage.getItem('access_token')) {
      this.authorized = true;
    }
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

  public register(username: string, password: string, token: string): Observable<void> {
    return this.post('register', {username: username, password: password, token: token}).map(res => {
      if (res.hasOwnProperty('token')) {
        localStorage.setItem('access_token', res.token);
        this.authorized = true;
      }
    });
  }

  public get(url: string): Observable<any> {
    return this.request('get', url);
  }

  public post(url: string, data: any): Observable<any> {
    return this.request('post', url, data);
  }

  public patch(url: string, data: any): Observable<any> {
    return this.request('patch', url, data);
  }

  public request(requestMethod: string, url: string, data?: any): Observable<any> {
    return Observable.create(observer => {
      if (!url.startsWith('/')) {
        url = '/' + url;
      }

      url = 'http://localhost:8080/api/v1' + url;

      return this.http[requestMethod](url, data).subscribe(result => {
        observer.next(result);
        observer.complete();
      }, error => {
          if (error.status === 401) {
            localStorage.removeItem('access_token');
            this.authorized = false;
            this.router.navigateByUrl('/login');
          }

          this.messageService.error(error.error.error_message);

          observer.error(error);
        });
    });

    // if (!url.startsWith('/')) {
    //   url = '/' + url;
    // }
    //
    // url = 'http://localhost:8080/api/v1' + url;
    //
    // return this.http[requestMethod](url, data).catch(error => {
    //   this.handleError<any[]>(requestMethod, []);
    //   return Observable.throw(error);
    // });

    // this.messageService.error("aaaaaaaaa");
    // return this.http.post(url, data).pipe(
    //   // tap(_ => console.log("")),
    //   catchError(this.handleError<any[]>('post', []))
    // );
  }
  //
  // private handleError<T> (operation = 'operation', result?: T) {
  //   this.messageService.error('shit happens');
  //   return (error: any): Observable<T> => {
  //     this.messageService.error('shit happens1');
  //
  //     // TODO: send the error to remote logging infrastructure
  //     // console.error(error); // log to console instead
  //     // TODO: better job of transforming error for user consumption
  //     // this.log(`${operation} failed: ${error.message}`);
  //
  //     // Let the app keep running by returning an empty result.
  //     return of(result as T);
  //   };
  // }
}
