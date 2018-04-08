import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyWordsPageComponent } from './my-words-page.component';

describe('MyWordsPageComponent', () => {
  let component: MyWordsPageComponent;
  let fixture: ComponentFixture<MyWordsPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyWordsPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyWordsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
