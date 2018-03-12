import {Component, OnInit, ViewChild} from '@angular/core';
import {AuthService} from '../auth.service';
import {Router} from '@angular/router';
import {FormGroup} from '@angular/forms';

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
    private authService: AuthService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  login() {
    this.authService.login(this.username, this.password).subscribe(() => {
      this.router.navigateByUrl('/');
    }, error => {
      this.errorString = error.error.message;
    });
  }
}
