import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-audio-play-button',
  templateUrl: './audio-play-button.component.html',
  styleUrls: ['./audio-play-button.component.css']
})
export class AudioPlayButtonComponent implements OnInit {
  @Input() url: string;
  @Input() title: string;

  public audio: any = new Audio();
  public isPlaying = false;
  /*  public get isPlaying() {
    return !this.audio.paused && this.audio.currentTime;
  }*/

  constructor() { }

  ngOnInit() {
    this.audio.src = this.url;
    this.audio.load();
    this.audio.onplaying = () => this.isPlaying = true;
    this.audio.onended = () => this.isPlaying = false;
    this.audio.onpause = () => this.isPlaying = false;
  }

  toggle() {
    if (this.isPlaying)
      this.stop();
    else
      this.play();
  }

  play() {
    this.audio.play();
  }
  stop() {
    this.audio.pause();
    this.audio.currentTime = 0;
  }

}
