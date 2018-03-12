import { Injectable } from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {HttpClient} from '@angular/common/http';
import {catchError} from 'rxjs/operators';
import {of} from 'rxjs/observable/of';

@Injectable()
export class ApiService {

  constructor(private http: HttpClient) { }

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

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      // this.log(`${operation} failed: ${error.message}`);

      // this.messageService.error(error.message);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
