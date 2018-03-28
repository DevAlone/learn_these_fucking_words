import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AudioPlayButtonComponent } from './audio-play-button.component';

describe('AudioPlayButtonComponent', () => {
  let component: AudioPlayButtonComponent;
  let fixture: ComponentFixture<AudioPlayButtonComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AudioPlayButtonComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AudioPlayButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
