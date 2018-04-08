import { Component, OnInit } from '@angular/core';
import {ApiService} from "../api.service";

@Component({
  selector: 'app-debug-page',
  templateUrl: './debug-page.component.html',
  styleUrls: ['./debug-page.component.css']
})
export class DebugPageComponent implements OnInit {
  users: any[];
  languages: any[];
  words: any[];
  memorizations: any[];

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.api.get('users').subscribe(res => {
      this.users = res;
    });
    this.api.get('words').subscribe(res => {
      this.words = res;
    });
    this.api.get('languages').subscribe(res => {
      this.languages = res;
    });
    this.api.get('memorizations').subscribe(res => {
      this.memorizations = res;
    });
  }
}
