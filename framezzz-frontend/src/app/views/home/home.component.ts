import { Component, inject } from '@angular/core';
import { ComicsCardComponent } from '../../components/comics-card/comics-card.component';
import { ToolbarComponent } from '../../components/toolbar/toolbar.component';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { ComicsService } from '../../services/comics/comics.service';
import { Observable } from 'rxjs';
import { Comic } from '../../interfaces/comic.interface';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'app-home',
  imports: [
    ComicsCardComponent,
    ToolbarComponent,
    MatIconModule,
    MatButtonModule,
    AsyncPipe,
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.sass',
})
export class HomeComponent {
  private readonly comicsService = inject(ComicsService);

  private currentComicId: number = 1;

  protected comic$!: Observable<Comic>;

  public ngOnInit() {
    this.comic$ = this.comicsService.getComic(this.currentComicId);
  }

  protected previousCard() {
    if (this.currentComicId > 1) {
      this.comic$ = this.comicsService.getComic(this.currentComicId - 1);
      this.currentComicId -= 1;
    }
  }

  protected nextCard() {
    this.comic$ = this.comicsService.getComic(this.currentComicId + 1);
    this.currentComicId += 1;
  }
}
