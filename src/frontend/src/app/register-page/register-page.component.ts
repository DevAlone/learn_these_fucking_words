import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {ApiService} from "../api.service";

@Component({
  selector: 'app-register-page',
  templateUrl: './register-page.component.html',
  styleUrls: ['./register-page.component.css']
})
export class RegisterPageComponent implements OnInit {
  username: string;
  password: string;
  token: string;
  errorString: string;

  constructor(
    private api: ApiService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  register() {
    this.api.register(this.username, this.password, this.token).subscribe((res) => {
      console.log('ok');
      this.router.navigateByUrl('/');
    }, error => {
      console.log('error happened');
      console.log(error);
      if (error.hasOwnProperty('error')) {
        this.errorString = error.error.error_message;
      } else {
        this.errorString = error.message;
      }
    });
  }
}
