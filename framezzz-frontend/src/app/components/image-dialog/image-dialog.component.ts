import { Component, inject, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatDialogModule } from '@angular/material/dialog';

@Component({
  selector: 'app-image-dialog',
  imports: [MatDialogModule],
  templateUrl: './image-dialog.component.html',
  styleUrl: './image-dialog.component.sass',
})
export class ImageDialogComponent {
  dialogRef = inject(MatDialogRef<ImageDialogComponent>);

  data = inject(MAT_DIALOG_DATA);
}
