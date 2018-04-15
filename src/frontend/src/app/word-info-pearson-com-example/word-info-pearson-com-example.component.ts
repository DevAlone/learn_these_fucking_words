import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-word-info-pearson-com-example',
  templateUrl: './word-info-pearson-com-example.component.html',
  styleUrls: ['./word-info-pearson-com-example.component.css']
})
export class WordInfoPearsonComExampleComponent implements OnInit {
  @Input() exampleData: any;
  constructor() { }

  ngOnInit() {
  }

}
