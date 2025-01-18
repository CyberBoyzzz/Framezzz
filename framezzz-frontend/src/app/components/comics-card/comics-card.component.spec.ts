import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ComicsCardComponent } from './comics-card.component';
import { MatDialog } from '@angular/material/dialog';
import { NO_ERRORS_SCHEMA } from '@angular/core';

describe('ComicsCardComponent', () => {
  let component: ComicsCardComponent;
  let fixture: ComponentFixture<ComicsCardComponent>;

  const mockComic = {
    id: 1,
    title: 'Mock Comic Title',
    description: 'Mock description',
    img: 'https://example.com/mock-image.jpg',
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ComicsCardComponent],
      providers: [{ provide: MatDialog, useValue: {} }],
      schemas: [NO_ERRORS_SCHEMA],
    }).compileComponents();

    fixture = TestBed.createComponent(ComicsCardComponent);
    component = fixture.componentInstance;
    component.comic = mockComic;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
