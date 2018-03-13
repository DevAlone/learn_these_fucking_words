import { Component, OnInit } from '@angular/core';
import {ApiService} from "../api.service";

@Component({
  selector: 'app-my-words-page',
  templateUrl: './my-words-page.component.html',
  styleUrls: ['./my-words-page.component.css']
})
export class MyWordsPageComponent implements OnInit {
  memorizations: any;
  newWordLanguage: number = 1;
  newWord: string = "";
  constructor(private api: ApiService) { }

  ngOnInit() {
    this.api.get('my/memorizations').subscribe(result => {
      this.memorizations = result;
    });
  }

  addNewWord() {
    console.log(this.newWordLanguage);
    console.log(this.newWord);
    this.api.post('words', {word: this.newWord, languageId: this.newWordLanguage}).subscribe(result => {
      // TODO: push to the list
      console.log(result);
    });
  }
}
