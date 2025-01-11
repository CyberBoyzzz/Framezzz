import { Component } from '@angular/core';
import { ComicsCardComponent } from '../../components/comics-card/comics-card.component';

@Component({
  selector: 'app-home',
  imports: [ComicsCardComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.sass'
})
export class HomeComponent {

}
