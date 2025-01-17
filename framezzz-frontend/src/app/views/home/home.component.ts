import { Component, inject } from '@angular/core';
import { ComicsCardComponent } from '../../components/comics-card/comics-card.component';
import { ToolbarComponent } from '../../components/toolbar/toolbar.component';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { ComicsService } from '../../services/comics/comics.service';
import { Observable, Subscription } from 'rxjs';
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

  private subscriptions: Subscription[] = [];

  private currentComicId: number = 1;

  protected currentComic: Comic = {
    title: '',
    id: 0,
    img: '',
  };

  protected comic$!: Observable<Comic>;

  public ngOnInit(): void {
    this.subscriptions.push(
      this.comicsService.getComic(this.currentComicId).subscribe({
        next: (response: Comic) => {
          this.currentComic = response;
        },
      })
    );
  }

  protected previousCard(): void {
    if (this.currentComicId > 1) {
      this.subscriptions.push(
        this.comicsService.getComic(this.currentComicId - 1).subscribe({
          next: (response: Comic) => {
            this.currentComic = response;
          },
        })
      );

      this.currentComicId -= 1;
    }
  }

  protected nextCard(): void {
    this.subscriptions.push(
      this.comicsService.getComic(this.currentComicId + 1).subscribe({
        next: (response: Comic) => {
          this.currentComic = response;
        },
      })
    );

    this.currentComicId += 1;
  }

  public ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription?.unsubscribe());
  }
}
