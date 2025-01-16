import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-comics-card',
  imports: [MatCardModule, MatButtonModule, MatIconModule],
  templateUrl: './comics-card.component.html',
  styleUrl: './comics-card.component.sass',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ComicsCardComponent {
  @Input() protected image: string = '';

  @Input() protected title: string = '';

  @Input() protected likesCount: number = 0;
}
