import { Component, OnInit, Input } from '@angular/core';
import { ApiService } from '../api.service';
import { Word } from '../models/word';

@Component({
  selector: 'app-word-info-pearson-com',
  templateUrl: './word-info-pearson-com.component.html',
  styleUrls: ['./word-info-pearson-com.component.css']
})
export class WordInfoPearsonComComponent implements OnInit {
  @Input() word: Word;
  wordInformations: any[];

  constructor(private api: ApiService) { }

  ngOnInit() {
    var url = '/word_info_items/pearson.com/' + this.word.word;
    this.api.get(url).subscribe(res => {
      this.wordInformations = res.data;
    })
  }
}
