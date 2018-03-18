import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-learning-page',
  templateUrl: './learning-page.component.html',
  styleUrls: ['./learning-page.component.css']
})
export class LearningPageComponent implements OnInit {
  memorization: any;
  word: any;
  sent = false;

  constructor(
    private api: ApiService,
    private router: Router
  ) { }

  ngOnInit() {
      this.nextWord();
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
    }, error => {
      if (error.status === 404) {
        this.router.navigateByUrl('/my/words');
      }
    });
  }
}
