import {Component, OnInit, ViewChild} from '@angular/core';
import {Router} from '@angular/router';
import {FormGroup} from '@angular/forms';
import {ApiService} from "../api.service";

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {
  // @ViewChild('loginForm') loginForm;
  username: string;
  password: string;
  errorString: string;

  constructor(
    private api: ApiService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  login() {
    this.api.login(this.username, this.password).subscribe((res) => {
      this.router.navigateByUrl('/');
    }, error => {
      if (error.hasOwnProperty('error'))
        this.errorString = error.error.error_message;
      else
        this.errorString = error.message;
    });
  }
}
