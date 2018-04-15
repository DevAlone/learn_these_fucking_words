import {
  Component, OnInit, ViewChild, ElementRef, QueryList
} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';
import { Word } from '../models/word';
import { MatButtonToggleGroup } from '@angular/material';

@Component({
  selector: 'app-learning-page',
  templateUrl: './learning-page.component.html',
  styleUrls: ['./learning-page.component.css']
})
export class LearningPageComponent implements OnInit {
  memorization: any;
  word: Word;
  sent = false;
  chosenKnowledge: number = 0;

  constructor(
    private api: ApiService,
    private router: Router,
    private http: HttpClient
  ) {
  }

  ngOnInit() {
    this.nextWord();
  }
  ngAfterViewInit(): void {
  }

  sendKnowledge(value: number) {
    this.memorization.memorizationCoefficient = value / 4.0;

    const data = {
      memorizationCoefficient: this.memorization.memorizationCoefficient
    };

    this.api.patch('my/memorizations/' + this.word.id, data).subscribe(result => {
      this.memorization = result.data;
      this.sent = true;
    });
  }

  nextWord() {
    this.api.get('learning/word').subscribe(result => {
      this.memorization = result.data;
      this.word = this.memorization.word;
      this.sent = false;
      this.chosenKnowledge = (this.memorization.memorizationCoefficient * 4) | 0;
    }, error => {
      if (error.status === 404) {
        this.router.navigateByUrl('/my/words');
      }
    });
  }
}
