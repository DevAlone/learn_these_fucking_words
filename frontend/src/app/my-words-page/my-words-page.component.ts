import { Component, OnInit } from '@angular/core';
import {ApiService} from "../api.service";

@Component({
  selector: 'app-my-words-page',
  templateUrl: './my-words-page.component.html',
  styleUrls: ['./my-words-page.component.css']
})
export class MyWordsPageComponent implements OnInit {
  memorizations: any;
  newWordLanguage: any;
  newWord: string = ""; 
  // TODO: load from server
  languages: any[] = [
  {code: 'eng'},
  {code: 'rus'},
  ];

  constructor(private api: ApiService) {
    this.newWordLanguage = this.languages[0];
  }

  ngOnInit() {
    this.api.get('my/memorizations').subscribe(result => {
      this.memorizations = result;
    });
  }

  addNewWord() {
    console.log(this.newWordLanguage);
    console.log(this.newWord);
    var data = {
        word: this.newWord, 
        languageCode: this.newWordLanguage.code
    };

    this.api.post('words', data).subscribe(result => {
      // TODO: push to the list
      console.log(result);
    });
  }
}
