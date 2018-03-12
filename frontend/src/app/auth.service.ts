import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map'

@Injectable()
export class AuthService {

  constructor(private http: HttpClient) { }

  login(username: string, password: string): Observable<void> {
    return this.http.post<any>('/api/v1/login', {username: username, password: password}).map(res => {
      console.log(res);
      localStorage.setItem('access_token', res.token);
    });
  }
}
