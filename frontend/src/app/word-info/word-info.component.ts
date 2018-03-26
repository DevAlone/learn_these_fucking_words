import { Component, OnInit, Input } from '@angular/core';
import { Word } from '../models/word';
import { ApiService } from '../api.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-word-info',
  templateUrl: './word-info.component.html',
  styleUrls: ['./word-info.component.css']
})
export class WordInfoComponent implements OnInit {
  @Input() word: Word;
  images: any[];

  constructor(
    private api: ApiService,
    private http: HttpClient
  ) { }

  ngOnInit() {
    this.api.get('images/' + this.word.word).subscribe(result => {
      this.images = result.data;
    });
  }

}
