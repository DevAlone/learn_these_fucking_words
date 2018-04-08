import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordInfoPearsonComComponent } from './word-info-pearson-com.component';

describe('WordInfoPearsonComComponent', () => {
  let component: WordInfoPearsonComComponent;
  let fixture: ComponentFixture<WordInfoPearsonComComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordInfoPearsonComComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordInfoPearsonComComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
