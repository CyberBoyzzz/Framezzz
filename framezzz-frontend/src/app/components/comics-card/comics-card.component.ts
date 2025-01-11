import { ChangeDetectionStrategy, Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';

@Component({
  selector: 'app-comics-card',
  imports: [MatCardModule, MatButtonModule],
  templateUrl: './comics-card.component.html',
  styleUrl: './comics-card.component.sass',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ComicsCardComponent {

}
