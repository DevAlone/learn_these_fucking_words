import { Component, OnInit } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ApiService} from '../api.service';

@Component({
  selector: 'app-words-page',
  templateUrl: './words-page.component.html',
  styleUrls: ['./words-page.component.css']
})
export class WordsPageComponent implements OnInit {
  words: any[];

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.api.get('words').subscribe(res => {
      this.words = res;
    });
  }
}
