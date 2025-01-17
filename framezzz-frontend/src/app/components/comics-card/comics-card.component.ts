import {
  ChangeDetectionStrategy,
  Component,
  inject,
  Input,
  SimpleChanges,
} from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { Comic } from '../../interfaces/comic.interface';
import { MatDialog } from '@angular/material/dialog';
import { ImageDialogComponent } from '../image-dialog/image-dialog.component';

@Component({
  selector: 'app-comics-card',
  imports: [MatCardModule, MatButtonModule, MatIconModule],
  templateUrl: './comics-card.component.html',
  styleUrl: './comics-card.component.sass',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ComicsCardComponent {
  @Input() public comic!: Comic;

  private readonly matDialog = inject(MatDialog);

  protected openImageDialog(img: string): void {
    this.matDialog.open(ImageDialogComponent, {
      data: { img },
      width: 'fit-content',
      height: '600px',
      maxWidth: '100%',
    });
  }
}
