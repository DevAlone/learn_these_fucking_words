import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordInfoPearsonComExampleComponent } from './word-info-pearson-com-example.component';

describe('WordInfoPearsonComExampleComponent', () => {
  let component: WordInfoPearsonComExampleComponent;
  let fixture: ComponentFixture<WordInfoPearsonComExampleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordInfoPearsonComExampleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordInfoPearsonComExampleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
