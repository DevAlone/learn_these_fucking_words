import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
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
  @ViewChild('tabs', { read: ElementRef }) tabs: ElementRef;
  @ViewChild('wikipediaFrame', { read: ElementRef }) wikipediaFrame: ElementRef;
  images: any[];

  constructor(
    private api: ApiService,
    private http: HttpClient,
    private sanitizer: DomSanitizer
  ) { }

  ngOnInit() {
    this.api.get('images/' + this.word.word).subscribe(result => {
      this.images = result.data;
    });
  }

  tabChanged(event: any) {
    this.tabs.nativeElement.scrollIntoView();
  }
}
